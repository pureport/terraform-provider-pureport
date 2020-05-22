package pureport

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/pureport/pureport-sdk-go/pureport/client"
	"github.com/terraform-providers/terraform-provider-pureport/pureport/configuration"
)

func init() {
	resource.AddTestSweepers("pureport_oracle_connection", &resource.Sweeper{
		Name: "pureport_oracle_connection",
		F: func(region string) error {
			c, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}

			config := c.(*configuration.Config)
			connections, err := config.GetAccConnections()
			if err != nil {
				return fmt.Errorf("Error getting connections %s", err)
			}

			if err = config.SweepConnections(connections); err != nil {
				return fmt.Errorf("Error occurred sweeping connections")
			}

			return nil
		},
	})
}

const testAccResourceOracleConnectionConfig_common = `

data "pureport_accounts" "main" {
  filter {
    name = "Name"
    values = ["Terraform"]
  }
}

data "pureport_locations" "main" {
  filter {
    name = "Name"
    values = ["^Washington*"]
  }
}

data "pureport_networks" "main" {
  account_href = data.pureport_accounts.main.accounts.0.href
  filter {
    name = "Name"
    values = ["Bansh.*"]
  }
}

data "pureport_cloud_regions" "main" {
  filter {
    name = "DisplayName"
    values = ["Ashburn"]
  }
}

`

func testAccResourceOracleConnectionConfig_basic() string {
	format := testAccResourceOracleConnectionConfig_common + `
resource "pureport_oracle_connection" "basic" {
  name = "%s"
  speed = "1000"

  location_href = data.pureport_locations.main.locations.0.href
  network_href = data.pureport_networks.main.networks.0.href
  cloud_region_href = data.pureport_cloud_regions.main.regions.0.href

  primary_ocid = "%s"
  secondary_ocid = "%s"

  bgp_peering {
    availability_domain = "PRIMARY"
    pureport_subnet = "169.254.16.1/30"
    customer_subnet = "169.254.16.2/30"
  }

  bgp_peering {
    availability_domain = "SECONDARY"
    pureport_subnet = "169.254.16.5/30"
    customer_subnet = "169.254.16.6/30"
  }

  tags = {
    Environment = "tf-test"
    sweep       = "TRUE"
  }
}
`

	connection_name := acctest.RandomWithPrefix("OracleFastConnectTest")

	return fmt.Sprintf(
		format,
		connection_name,
		testOraclePrimaryOCID,
		testOracleSecondaryOCID,
	)
}

func testAccResourceOracleConnectionConfig_basic_update_no_respawn() string {
	format := testAccResourceOracleConnectionConfig_common + `
resource "pureport_oracle_connection" "basic" {
  name = "%s"
  description = "Oracle Basic Test"
  speed = "1000"

  location_href = data.pureport_locations.main.locations.0.href
  network_href = data.pureport_networks.main.networks.0.href
  cloud_region_href = data.pureport_cloud_regions.main.regions.0.href

  primary_ocid = "%s"
  secondary_ocid = "%s"

  bgp_peering {
    availability_domain = "PRIMARY"
    pureport_subnet = "169.254.16.1/30"
    customer_subnet = "169.254.16.2/30"
  }

  bgp_peering {
    availability_domain = "SECONDARY"
    pureport_subnet = "169.254.16.5/30"
    customer_subnet = "169.254.16.6/30"
  }

  tags = {
    Environment = "tf-test"
    Owner       = "scott-pilgram"
    sweep       = "TRUE"
  }
}
`

	connection_name := acctest.RandomWithPrefix("OracleFastConnectTest")

	return fmt.Sprintf(
		format,
		connection_name,
		testOraclePrimaryOCID,
		testOracleSecondaryOCID,
	)
}

func testAccResourceOracleConnectionConfig_basic_update_respawn() string {
	format := testAccResourceOracleConnectionConfig_common + `
resource "pureport_oracle_connection" "basic" {
  name = "%s"
  description = "Oracle Basic Test"
  speed = "2000"

  location_href = data.pureport_locations.main.locations.0.href
  network_href = data.pureport_networks.main.networks.0.href
  cloud_region_href = data.pureport_cloud_regions.main.regions.0.href

  primary_ocid = "%s"
  secondary_ocid = "%s"

  bgp_peering {
    availability_domain = "PRIMARY"
    pureport_subnet = "169.254.16.1/30"
    customer_subnet = "169.254.16.2/30"
  }

  bgp_peering {
    availability_domain = "SECONDARY"
    pureport_subnet = "169.254.16.5/30"
    customer_subnet = "169.254.16.6/30"
  }
}
`
	connection_name := acctest.RandomWithPrefix("OracleFastConnectTest")

	return fmt.Sprintf(
		format,
		connection_name,
		testOraclePrimaryOCID,
		testOracleSecondaryOCID,
	)
}

func testAccResourceOracleConnectionConfig_invalid_ha() string {
	format := testAccResourceOracleConnectionConfig_common + `
resource "pureport_oracle_connection" "basic" {
  name = "%s"
  speed = "1000"
  high_availability = false

  location_href = "location/blah"
  network_href = "network/blah"
  cloud_region_href = data.pureport_cloud_regions.main.regions.0.href

  primary_ocid = "%s"
  secondary_ocid = "%s"

  bgp_peering {
    availability_domain = "PRIMARY"
    pureport_subnet = "169.254.16.1/30"
    customer_subnet = "169.254.16.2/30"
  }

  bgp_peering {
    availability_domain = "SECONDARY"
    pureport_subnet = "169.254.16.5/30"
    customer_subnet = "169.254.16.6/30"
  }

  tags = {
    Environment = "tf-test"
    Owner       = "scott-pilgram"
    sweep       = "TRUE"
  }
}
`

	connection_name := acctest.RandomWithPrefix("OracleFastConnectTest")

	return fmt.Sprintf(
		format,
		connection_name,
		testOraclePrimaryOCID,
		testOracleSecondaryOCID,
	)
}

func TestResourceOracleConnection_basic(t *testing.T) {

	if testEnvironmentName != "Production" {
		return
	}

	resourceName := "pureport_oracle_connection.basic"
	var instance client.OracleFastConnectConnection
	var respawn_instance client.OracleFastConnectConnection

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckOracleConnectionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceOracleConnectionConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckResourceOracleConnection(resourceName, &instance),

					resource.TestCheckResourceAttrPtr(resourceName, "id", &instance.Id),
					resource.TestMatchResourceAttr(resourceName, "name", regexp.MustCompile("^OracleFastConnectTest-.*")),
					resource.TestCheckResourceAttr(resourceName, "description", ""),
					resource.TestCheckResourceAttr(resourceName, "speed", "1000"),
					resource.TestCheckResourceAttr(resourceName, "high_availability", "true"),
					resource.TestCheckResourceAttr(resourceName, "location_href", "/locations/us-sea"),
					resource.TestMatchResourceAttr(resourceName, "network_href", regexp.MustCompile("/networks/network-.{16}")),
					resource.TestMatchResourceAttr(resourceName, "cloud_region_href", regexp.MustCompile("/cloudregions/oracle-.+")),
					resource.TestMatchResourceAttr(resourceName, "primary_ocid", regexp.MustCompile("ocid1.virtualcircuit.oc1.iad.[a-z0-9]+")),
					resource.TestMatchResourceAttr(resourceName, "secondary_ocid", regexp.MustCompile("ocid1.virtualcircuit.oc1.iad.[a-z0-9]+")),

					resource.TestCheckResourceAttr(resourceName, "gateways.#", "2"),

					resource.TestCheckResourceAttr(resourceName, "gateways.0.availability_domain", "PRIMARY"),
					resource.TestMatchResourceAttr(resourceName, "gateways.0.name", regexp.MustCompile("^OracleFastConnectTest-.* - Primary")),
					resource.TestCheckResourceAttr(resourceName, "gateways.0.description", ""),
					resource.TestCheckResourceAttr(resourceName, "gateways.0.customer_asn", "64512"),
					resource.TestMatchResourceAttr(resourceName, "gateways.0.customer_ip", regexp.MustCompile("169.254.[0-9]{1,3}.[0-9]{1,3}/30")),
					resource.TestCheckResourceAttr(resourceName, "gateways.0.pureport_asn", "394351"),
					resource.TestMatchResourceAttr(resourceName, "gateways.0.pureport_ip", regexp.MustCompile("169.254.[0-9]{1,3}.[0-9]{1,3}/30")),
					resource.TestCheckResourceAttrSet(resourceName, "gateways.0.bgp_password"),
					resource.TestMatchResourceAttr(resourceName, "gateways.0.peering_subnet", regexp.MustCompile("169.254.[0-9]{1,3}.[0-9]{1,3}")),
					resource.TestCheckResourceAttr(resourceName, "gateways.0.public_nat_ip", ""),
					resource.TestCheckResourceAttrSet(resourceName, "gateways.0.vlan"),
					resource.TestCheckResourceAttrSet(resourceName, "gateways.0.remote_id"),

					resource.TestCheckResourceAttr(resourceName, "gateways.1.availability_domain", "SECONDARY"),
					resource.TestMatchResourceAttr(resourceName, "gateways.1.name", regexp.MustCompile("^OracleFastConnectTest-.* - Secondary")),
					resource.TestCheckResourceAttr(resourceName, "gateways.1.description", ""),
					resource.TestCheckResourceAttr(resourceName, "gateways.1.customer_asn", "64512"),
					resource.TestMatchResourceAttr(resourceName, "gateways.1.customer_ip", regexp.MustCompile("169.254.[0-9]{1,3}.[0-9]{1,3}/30")),
					resource.TestCheckResourceAttr(resourceName, "gateways.1.pureport_asn", "394351"),
					resource.TestMatchResourceAttr(resourceName, "gateways.1.pureport_ip", regexp.MustCompile("169.254.[0-9]{1,3}.[0-9]{1,3}/30")),
					resource.TestCheckResourceAttrSet(resourceName, "gateways.1.bgp_password"),
					resource.TestMatchResourceAttr(resourceName, "gateways.1.peering_subnet", regexp.MustCompile("169.254.[0-9]{1,3}.[0-9]{1,3}")),
					resource.TestCheckResourceAttr(resourceName, "gateways.1.public_nat_ip", ""),
					resource.TestCheckResourceAttrSet(resourceName, "gateways.1.vlan"),
					resource.TestCheckResourceAttrSet(resourceName, "gateways.1.remote_id"),

					resource.TestCheckResourceAttr(resourceName, "nat_config.0.enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "nat_config.0.blocks.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "nat_config.0.mappings.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "nat_config.0.pnat_cidr", ""),

					resource.TestCheckResourceAttr(resourceName, "tags.Environment", "tf-test"),
					resource.TestCheckResourceAttr(resourceName, "tags.Owner", "scott-pilgram"),
				),
			},
			{
				Config: testAccResourceOracleConnectionConfig_basic_update_no_respawn(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPtr(resourceName, "id", &instance.Id),
					resource.TestMatchResourceAttr(resourceName, "name", regexp.MustCompile("^OracleFastConnectTest-.*")),
					resource.TestCheckResourceAttr(resourceName, "description", "Oracle Basic Test"),

					resource.TestCheckResourceAttr(resourceName, "speed", "1000"),
					resource.TestCheckResourceAttr(resourceName, "high_availability", "true"),

					resource.TestCheckResourceAttr(resourceName, "tags.Owner", "scott-pilgram"),
				),
			},
			{
				Config: testAccResourceOracleConnectionConfig_basic_update_respawn(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckResourceOracleConnection(resourceName, &respawn_instance),
					resource.TestCheckResourceAttrPtr(resourceName, "id", &respawn_instance.Id),
					TestCheckResourceConnectionIdChanged(&instance.Id, &respawn_instance.Id),
					resource.TestMatchResourceAttr(resourceName, "name", regexp.MustCompile("^OracleFastConnectTest-.*")),
					resource.TestCheckResourceAttr(resourceName, "description", "Oracle Basic Test"),
					resource.TestMatchResourceAttr(resourceName, "aws_account_id", regexp.MustCompile("[0-9]{12}")),
					resource.TestCheckResourceAttr(resourceName, "speed", "2000"),
					resource.TestCheckResourceAttr(resourceName, "high_availability", "true"),
				),
			},
		},
	})
}

func TestResourceOracleConnection_invalid_ha(t *testing.T) {
	resource.Test(t, resource.TestCase{
		IsUnitTest: true,
		PreCheck:   func() { testAccPreCheck(t) },
		Providers:  testAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      testAccResourceOracleConnectionConfig_invalid_ha(),
				ExpectError: regexp.MustCompile("Oracle Cloud Connection high availability required."),
			},
		},
	})
}

func testAccCheckResourceOracleConnection(name string, instance *client.OracleFastConnectConnection) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		config, ok := testAccProvider.Meta().(*configuration.Config)
		if !ok {
			return fmt.Errorf("Error getting Pureport client")
		}

		// Find the state object
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Can't find Oracle Connection resource: %s", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set")
		}

		id := rs.Primary.ID

		ctx := config.Session.GetSessionContext()
		found, resp, err := config.Session.Client.ConnectionsApi.GetConnection(ctx, id)

		if err != nil {
			return fmt.Errorf("receive error when requesting Oracle Connection %s", id)
		}

		if resp.StatusCode != 200 {
			return fmt.Errorf("Error getting Oracle Connection ID %s: %s", id, err)
		}

		*instance = found.(client.OracleFastConnectConnection)

		return nil
	}
}

func testAccCheckOracleConnectionDestroy(s *terraform.State) error {

	config, ok := testAccProvider.Meta().(*configuration.Config)
	if !ok {
		return fmt.Errorf("Error getting Pureport client")
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "pureport_oracle_connection" {
			continue
		}

		id := rs.Primary.ID

		ctx := config.Session.GetSessionContext()
		_, resp, err := config.Session.Client.ConnectionsApi.GetConnection(ctx, id)

		if err != nil && resp.StatusCode != 404 {
			return fmt.Errorf("should not get error for Oracle Connection with ID %s after delete: %s", id, err)
		}
	}

	return nil
}
