provider "azurerm" {
  version = "~> 1.44"
}

data "pureport_accounts" "main" {
  filter {
    name   = "Name"
    values = ["Terraform .*"]
  }
}

data "pureport_locations" "main" {
  filter {
    name   = "Name"
    values = ["Sea.*"]
  }
}

data "pureport_networks" "main" {
  account_href = data.pureport_accounts.main.accounts.0.href
  filter {
    name   = "Name"
    values = ["Bansh.*"]
  }
}

data "azurerm_express_route_circuit" "main" {
  name                = "terraform-acc-express-route-dev1"
  resource_group_name = "terraform-acceptance-tests"
}

resource "pureport_azure_connection" "main" {
  name              = "AzureExpressRouteTest-ksk"
  description       = "Some random description"
  speed             = "100"
  high_availability = true

  location_href = data.pureport_locations.main.locations.0.href
  network_href  = data.pureport_networks.main.networks.0.href

  service_key = data.azurerm_express_route_circuit.main.service_key

  tags = {
    Environment = "tf-test"
    Owner       = "ksk-azure"
    sweep       = "TRUE"
  }
}
