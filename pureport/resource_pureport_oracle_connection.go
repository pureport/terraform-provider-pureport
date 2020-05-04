package pureport

import (
	"fmt"
	"log"
	"net/url"
	"path/filepath"
	"regexp"

	"github.com/antihax/optional"
	"github.com/hashicorp/terraform-plugin-sdk/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/structure"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/pureport/pureport-sdk-go/pureport/client"
	"github.com/terraform-providers/terraform-provider-pureport/pureport/configuration"
	"github.com/terraform-providers/terraform-provider-pureport/pureport/connection"
	"github.com/terraform-providers/terraform-provider-pureport/pureport/tags"
)

var (
	OracleOCIDSchema = schema.Schema{
		Type:         schema.TypeString,
		Required:     true,
		ForceNew:     true,
		ValidateFunc: validation.StringMatch(regexp.MustCompile(`ocid1.virtualcircuit.oc1.iad.[a-z0-9]{60}`), "Must be a valid Oracle Cloud Virtual Circuit ID."),
	}
)

func OracleResourceDiff(d *schema.ResourceDiff) error {
	highAvailability := false
	if v, ok := d.GetOk("high_availability"); ok {
		highAvailability = v.(bool)
	}

	if !highAvailability {
		return fmt.Errorf("Oracle Cloud Connection high availability required.")
	}

	return nil
}

func resourceOracleConnection() *schema.Resource {

	connection_schema := map[string]*schema.Schema{
		"primary_ocid":   &OracleOCIDSchema,
		"secondary_ocid": &OracleOCIDSchema,
		"bgp_peering": {
			Type:     schema.TypeList,
			Required: true,
			MaxItems: 2,
			MinItems: 2,
			Elem: &schema.Resource{
				Schema: connection.BGPPeerSchema,
			},
		},
		"cloud_region_href": {
			Type:     schema.TypeString,
			Required: true,
		},
		"speed": {
			Type:         schema.TypeInt,
			Required:     true,
			ForceNew:     true,
			ValidateFunc: validation.IntInSlice([]int{1000, 2000}),
		},
		"gateways": {
			Computed: true,
			Type:     schema.TypeList,
			MinItems: 1,
			MaxItems: 2,
			Elem: &schema.Resource{
				Schema: connection.StandardGatewaySchema,
			},
		},
	}

	// Add the base items
	for k, v := range connection.GetBaseResourceConnectionSchema() {
		connection_schema[k] = v
	}

	return &schema.Resource{
		Create: resourceOracleConnectionCreate,
		Read:   resourceOracleConnectionRead,
		Update: resourceOracleConnectionUpdate,
		Delete: resourceOracleConnectionDelete,

		Schema: connection_schema,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(connection.CreateTimeout),
			Delete: schema.DefaultTimeout(connection.DeleteTimeout),
		},
		CustomizeDiff: customdiff.If(
			customdiff.ResourceConditionFunc(func(d *schema.ResourceDiff, meta interface{}) bool {
				return d.HasChange("high_availability")
			}),
			schema.CustomizeDiffFunc(func(d *schema.ResourceDiff, meta interface{}) error {
				return OracleResourceDiff(d)
			}),
		),
	}
}

func expandBGPPeering(d *schema.ResourceData) *client.PeeringConfiguration {

	pc := &client.PeeringConfiguration{
		Type: "PRIVATE",
	}

	// Get the BGP peering data
	bgpPeering := d.Get("bgp_peering").([]interface{})

	for _, v := range bgpPeering {
		config := v.(map[string]interface{})

		switch config["availability_domain"].(string) {
		case "PRIMARY":
			pc.PrimaryPureportBgpIP = config["pureport_subnet"].(string)
			pc.PrimaryRemoteBgpIP = config["customer_subnet"].(string)
		case "SECONDARY":
			pc.SecondaryPureportBgpIP = config["pureport_subnet"].(string)
			pc.SecondaryRemoteBgpIP = config["customer_subnet"].(string)
		}
	}

	return pc
}

func expandOracleConnection(d *schema.ResourceData) client.OracleFastConnectConnection {

	// Generic Connection values
	speed := d.Get("speed").(int)

	// Create the body of the request
	c := client.OracleFastConnectConnection{
		Type:  "ORACLE_FAST_CONNECT",
		Name:  d.Get("name").(string),
		Speed: int32(speed),
		Location: client.Link{
			Href: d.Get("location_href").(string),
		},
		Network: client.Link{
			Href: d.Get("network_href").(string),
		},
		CloudRegion: client.Link{
			Href: d.Get("cloud_region_href").(string),
		},
		PrimaryOcid:      d.Get("primary_ocid").(string),
		SecondaryOcid:    d.Get("secondary_ocid").(string),
		BillingTerm:      d.Get("billing_term").(string),
		HighAvailability: true,
	}

	c.CustomerNetworks = connection.ExpandCustomerNetworks(d)
	c.Nat = connection.ExpandNATConfiguration(d)
	c.Peering = expandBGPPeering(d)

	if description, ok := d.GetOk("description"); ok {
		c.Description = description.(string)
	}

	if highAvailability, ok := d.GetOk("high_availability"); ok {
		c.HighAvailability = highAvailability.(bool)
	}

	if t, ok := d.GetOk("tags"); ok {
		c.Tags = tags.FilterTags(t.(map[string]interface{}))
	}

	return c
}

func resourceOracleConnectionCreate(d *schema.ResourceData, m interface{}) error {

	c := expandOracleConnection(d)

	config := m.(*configuration.Config)
	ctx := config.Session.GetSessionContext()

	opts := client.AddConnectionOpts{
		Connection: optional.NewInterface(c),
	}

	_, resp, err := config.Session.Client.ConnectionsApi.AddConnection(
		ctx,
		filepath.Base(c.Network.Href),
		&opts,
	)

	if err != nil {

		http_err := err
		json_response := string(err.(client.GenericOpenAPIError).Body()[:])
		response, err := structure.ExpandJsonFromString(json_response)

		if err != nil {
			log.Printf("Error Creating new %s: %v", connection.OracleConnectionName, err)

		} else {
			statusCode := int(response["status"].(float64))

			log.Printf("Error Creating new %s: %d\n", connection.OracleConnectionName, statusCode)
			log.Printf("  %s\n", response["code"])
			log.Printf("  %s\n", response["message"])
		}

		d.SetId("")
		return fmt.Errorf("Error while creating %s: err=%s", connection.OracleConnectionName, http_err)
	}

	if resp.StatusCode >= 300 {
		d.SetId("")
		return fmt.Errorf("Error while creating %s: code=%v", connection.OracleConnectionName, resp.StatusCode)
	}

	loc := resp.Header.Get("location")
	u, err := url.Parse(loc)
	if err != nil {
		return fmt.Errorf("Error when decoding Connection ID")
	}

	id := filepath.Base(u.Path)
	d.SetId(id)

	if id == "" {
		log.Printf("Error when decoding location header")
		return fmt.Errorf("Error decoding Connection ID")
	}

	if err := connection.WaitForConnection(connection.OracleConnectionName, d, m); err != nil {
		return fmt.Errorf("Error waiting for %s: err=%s", connection.OracleConnectionName, err)
	}

	return resourceOracleConnectionRead(d, m)
}

func resourceOracleConnectionRead(d *schema.ResourceData, m interface{}) error {

	config := m.(*configuration.Config)
	connectionId := d.Id()
	ctx := config.Session.GetSessionContext()

	c, resp, err := config.Session.Client.ConnectionsApi.GetConnection(ctx, connectionId)
	if err != nil {
		if resp.StatusCode == 404 {
			log.Printf("Error Response while reading %s: code=%v", connection.OracleConnectionName, resp.StatusCode)
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error reading data for %s: %s", connection.OracleConnectionName, err)
	}

	if resp.StatusCode >= 300 {
		return fmt.Errorf("Error Response while reading %s: code=%v", connection.OracleConnectionName, resp.StatusCode)
	}

	conn := c.(client.OracleFastConnectConnection)
	d.Set("primary_ocid", conn.PrimaryOcid)
	d.Set("secondary_ocid", conn.SecondaryOcid)
	d.Set("description", conn.Description)
	d.Set("high_availability", conn.HighAvailability)
	d.Set("href", conn.Href)
	d.Set("name", conn.Name)
	d.Set("speed", conn.Speed)
	d.Set("state", conn.State)

	if err := d.Set("customer_networks", connection.FlattenCustomerNetworks(conn.CustomerNetworks)); err != nil {
		return fmt.Errorf("Error setting customer networks for %s %s: %s", connection.OracleConnectionName, d.Id(), err)
	}

	// NAT Configuration
	if err := d.Set("nat_config", connection.FlattenNatConfig(conn.Nat)); err != nil {
		return fmt.Errorf("Error setting NAT Configuration for %s %s: %s", connection.OracleConnectionName, d.Id(), err)
	}

	// Add Gateway information
	var gateways []map[string]interface{}
	if g := conn.PrimaryGateway; g != nil {
		gateways = append(gateways, connection.FlattenStandardGateway(g))
	}
	if g := conn.SecondaryGateway; g != nil {
		gateways = append(gateways, connection.FlattenStandardGateway(g))
	}
	if err := d.Set("gateways", gateways); err != nil {
		return fmt.Errorf("Error setting gateway information for %s %s: %s", connection.OracleConnectionName, d.Id(), err)
	}

	if err := d.Set("location_href", conn.Location.Href); err != nil {
		return fmt.Errorf("Error setting location for %s %s: %s", connection.OracleConnectionName, d.Id(), err)
	}

	if err := d.Set("network_href", conn.Network.Href); err != nil {
		return fmt.Errorf("Error setting network for %s %s: %s", connection.OracleConnectionName, d.Id(), err)
	}

	if err := d.Set("cloud_region_href", conn.CloudRegion.Href); err != nil {
		return fmt.Errorf("Error setting cloud region for %s %s: %s", connection.OracleConnectionName, d.Id(), err)
	}

	if err := d.Set("tags", conn.Tags); err != nil {
		return fmt.Errorf("Error setting tags for %s %s: %s", connection.OracleConnectionName, d.Id(), err)
	}

	return nil
}

func resourceOracleConnectionUpdate(d *schema.ResourceData, m interface{}) error {

	c := expandOracleConnection(d)

	d.Partial(true)

	config := m.(*configuration.Config)
	ctx := config.Session.GetSessionContext()

	if d.HasChange("name") {
		c.Name = d.Get("name").(string)
		d.SetPartial("name")
	}

	if d.HasChange("description") {
		c.Description = d.Get("description").(string)
		d.SetPartial("description")
	}

	if d.HasChange("speed") {
		c.Speed = int32(d.Get("speed").(int))
		d.SetPartial("speed")
	}

	if d.HasChange("customer_networks") {
		c.CustomerNetworks = connection.ExpandCustomerNetworks(d)
	}

	if d.HasChange("nat_config") {
		c.Nat = connection.ExpandNATConfiguration(d)
	}

	if d.HasChange("billing_term") {
		c.BillingTerm = d.Get("billing_term").(string)
	}

	if d.HasChange("tags") {
		_, nraw := d.GetChange("tags")
		c.Tags = tags.FilterTags(nraw.(map[string]interface{}))
	}

	opts := client.UpdateConnectionOpts{
		Connection: optional.NewInterface(c),
	}

	_, resp, err := config.Session.Client.ConnectionsApi.UpdateConnection(
		ctx,
		d.Id(),
		&opts,
	)

	if err != nil {

		if swerr, ok := err.(client.GenericOpenAPIError); ok {

			json_response := string(swerr.Body()[:])
			response, jerr := structure.ExpandJsonFromString(json_response)

			if jerr == nil {
				statusCode := int(response["status"].(float64))
				log.Printf("Error updating %s: %d\n", connection.OracleConnectionName, statusCode)
				log.Printf("  %s\n", response["code"])
				log.Printf("  %s\n", response["message"])
			}
		}

		return fmt.Errorf("Error while updating %s: err=%s", connection.OracleConnectionName, err)
	}

	if resp.StatusCode >= 300 {
		return fmt.Errorf("Error Response while updating %s: code=%v", connection.OracleConnectionName, resp.StatusCode)
	}

	if err := connection.WaitForConnection(connection.OracleConnectionName, d, m); err != nil {
		return fmt.Errorf("Error waiting for %s: err=%s", connection.OracleConnectionName, err)
	}

	d.Partial(false)

	return resourceOracleConnectionRead(d, m)
}

func resourceOracleConnectionDelete(d *schema.ResourceData, m interface{}) error {
	return connection.DeleteConnection(connection.OracleConnectionName, d, m)
}
