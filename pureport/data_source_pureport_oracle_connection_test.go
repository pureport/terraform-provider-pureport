package pureport

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

const testAccDataSourceOracleConnectionConfig_common = `
data "pureport_accounts" "main" {
  filter {
    name = "Name"
    values = ["Terraform .*"]
  }
}

data "pureport_networks" "main" {
  account_href = data.pureport_accounts.main.accounts.0.href
  filter {
    name = "Name"
    values = ["A Flock of Seagulls"]
  }
}

data "pureport_connections" "main" {
  network_href = data.pureport_networks.main.networks.0.href
  filter {
    name = "Name"
    values = ["ORACLE"]
  }
}
`

const testAccDataSourceOracleConnectionConfig_basic = testAccDataSourceOracleConnectionConfig_common + `
data "pureport_oracle_connection" "basic" {
  connection_id = data.pureport_connections.main.connections.0.id
}
`

func TestDataSourceOracleConnection_basic(t *testing.T) {

	if testEnvironmentName != "Production" {
		return
	}

	resourceName := "data.pureport_oracle_connection.basic"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceOracleConnectionConfig_basic,
				Check: resource.ComposeTestCheckFunc(

					resource.ComposeAggregateTestCheckFunc(

						resource.TestMatchResourceAttr(resourceName, "id", regexp.MustCompile("conn-.{16}")),
						resource.TestMatchResourceAttr(resourceName, "primary_ocid", regexp.MustCompile("ocid1.virtualcircuit.oc1.iad.[a-z0-9]{60}")),
						resource.TestMatchResourceAttr(resourceName, "secondary_ocid", regexp.MustCompile("ocid1.virtualcircuit.oc1.iad.[a-z0-9]{60}")),
						resource.TestCheckResourceAttr(resourceName, "speed", "1000"),
						resource.TestMatchResourceAttr(resourceName, "href", regexp.MustCompile("/connections/conn-.{16}")),
						resource.TestCheckResourceAttr(resourceName, "name", "ORACLE Connection DataSource"),
						resource.TestCheckResourceAttr(resourceName, "description", ""),
						resource.TestCheckResourceAttr(resourceName, "state", "ACTIVE"),
						resource.TestCheckResourceAttr(resourceName, "high_availability", "true"),
						resource.TestMatchResourceAttr(resourceName, "network_href", regexp.MustCompile("/networks/network-.{16}")),
						resource.TestCheckResourceAttr(resourceName, "location_href", "/locations/us-wdc"),

						resource.TestCheckResourceAttr(resourceName, "tags.#", "0"),

						resource.TestCheckResourceAttr(resourceName, "gateways.#", "2"),

						resource.TestCheckResourceAttr(resourceName, "gateways.0.availability_domain", "PRIMARY"),
						resource.TestCheckResourceAttr(resourceName, "gateways.0.name", "ORACLE Connection DataSource - Primary"),
						resource.TestCheckResourceAttr(resourceName, "gateways.0.description", ""),
						resource.TestCheckResourceAttr(resourceName, "gateways.0.customer_asn", "31898"),
						resource.TestMatchResourceAttr(resourceName, "gateways.0.customer_ip", regexp.MustCompile("169.254.[0-9]{1,3}.[0-9]{1,3}/30")),
						resource.TestCheckResourceAttr(resourceName, "gateways.0.pureport_asn", "394351"),
						resource.TestMatchResourceAttr(resourceName, "gateways.0.pureport_ip", regexp.MustCompile("169.254.[0-9]{1,3}.[0-9]{1,3}/30")),
						resource.TestMatchResourceAttr(resourceName, "gateways.0.peering_subnet", regexp.MustCompile("169.254.[0-9]{1,3}.[0-9]{1,3}")),
						resource.TestCheckResourceAttr(resourceName, "gateways.0.public_nat_ip", ""),
						resource.TestCheckResourceAttrSet(resourceName, "gateways.0.vlan"),
						resource.TestCheckResourceAttrSet(resourceName, "gateways.0.remote_id"),

						resource.TestCheckResourceAttr(resourceName, "gateways.1.availability_domain", "SECONDARY"),
						resource.TestCheckResourceAttr(resourceName, "gateways.1.name", "ORACLE Connection DataSource - Secondary"),
						resource.TestCheckResourceAttr(resourceName, "gateways.1.description", ""),
						resource.TestCheckResourceAttr(resourceName, "gateways.1.customer_asn", "31898"),
						resource.TestMatchResourceAttr(resourceName, "gateways.1.customer_ip", regexp.MustCompile("169.254.[0-9]{1,3}.[0-9]{1,3}/30")),
						resource.TestCheckResourceAttr(resourceName, "gateways.1.pureport_asn", "394351"),
						resource.TestMatchResourceAttr(resourceName, "gateways.1.pureport_ip", regexp.MustCompile("169.254.[0-9]{1,3}.[0-9]{1,3}/30")),
						resource.TestMatchResourceAttr(resourceName, "gateways.1.peering_subnet", regexp.MustCompile("169.254.[0-9]{1,3}.[0-9]{1,3}")),
						resource.TestCheckResourceAttr(resourceName, "gateways.1.public_nat_ip", ""),
						resource.TestCheckResourceAttrSet(resourceName, "gateways.1.vlan"),
						resource.TestCheckResourceAttrSet(resourceName, "gateways.1.remote_id"),
					),
				),
			},
		},
	})
}
