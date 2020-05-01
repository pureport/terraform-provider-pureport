package pureport

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/pureport/pureport-sdk-go/pureport/client"
	"github.com/terraform-providers/terraform-provider-pureport/pureport/configuration"
	"github.com/terraform-providers/terraform-provider-pureport/pureport/connection"
)

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

	if err := d.Set("tags", conn.Tags); err != nil {
		return fmt.Errorf("Error setting tags for %s %s: %s", connection.OracleConnectionName, d.Id(), err)
	}

	return nil
}
