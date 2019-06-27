data "pureport_cloud_regions" "us-west-2" {
  name_regex = "Oregon"
}

data "pureport_locations" "sea" {
  name_regex = "Seattle*"
}

data "aws_caller_identity" "current" {}

resource "pureport_aws_connection" "basic" {
  name              = "ConnectionTest-${count.index + 1}"
  speed             = "50"
  high_availability = false

  location_href = "${data.pureport_locations.sea.locations.0.href}"
  network_href  = "${var.connections_network_href}"

  aws_region     = "${data.pureport_cloud_regions.us-west-2.regions.0.identifier}"
  aws_account_id = "${data.aws_caller_identity.current.account_id}"

  count = 3
}

resource "pureport_aws_connection" "datasource" {
  name              = "AWS Connection DataSource"
  speed             = "50"
  high_availability = false

  location_href = "${data.pureport_locations.sea.locations.0.href}"
  network_href  = "${var.datasource_network_href}"

  aws_region     = "${data.pureport_cloud_regions.us-west-2.regions.0.identifier}"
  aws_account_id = "${data.aws_caller_identity.current.account_id}"
}

data "azurerm_express_route_circuit" "main" {
  name                = "terraform-acc-express-route-ds"
  resource_group_name = "terraform-acceptance-tests"
}

resource "pureport_azure_connection" "datasource" {
  name              = "AzureExpressRoute_DataSource"
  description       = "Some random description"
  speed             = "100"
  high_availability = true

  location_href = "${data.pureport_locations.sea.locations.0.href}"
  network_href  = "${var.datasource_network_href}"

  service_key = "${var.datasource_express_route_service_key}"
}

data "google_compute_network" "default" {
  name = "terraform-acc-network"
}

resource "google_compute_router" "main" {
  name    = "terraform-acc-ds-router-${count.index + 1}"
  network = "${data.google_compute_network.default.name}"

  bgp {
    asn = "16550"
  }

  count = 2
}

resource "google_compute_interconnect_attachment" "main" {
  name                     = "terraform-acc-ds-interconnect-${count.index + 1}"
  router                   = "${element(google_compute_router.main.*.self_link, count.index)}"
  type                     = "PARTNER"
  edge_availability_domain = "AVAILABILITY_DOMAIN_${count.index + 1}"

  lifecycle {
    ignore_changes = ["vlan_tag8021q"]
  }

  count = 2
}

resource "pureport_google_cloud_connection" "datasource" {
  name  = "GoogleCloud_DataSource"
  speed = "50"

  location_href = "${data.pureport_locations.sea.locations.0.href}"
  network_href  = "${var.datasource_network_href}"

  primary_pairing_key = "${google_compute_interconnect_attachment.main.0.pairing_key}"
}

resource "pureport_site_vpn_connection" "main" {
  name              = "SiteVPN_RouteBasedBGP_DataSource"
  speed             = "100"
  high_availability = true

  location_href = "${data.pureport_locations.sea.locations.0.href}"
  network_href  = "${var.datasource_network_href}"

  ike_version = "V2"

  routing_type = "ROUTE_BASED_BGP"
  customer_asn = 30000

  primary_customer_router_ip   = "123.123.123.123"
  secondary_customer_router_ip = "124.124.124.124"
}