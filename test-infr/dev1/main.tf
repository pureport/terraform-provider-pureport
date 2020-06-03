provider "azurerm" {
  version = "~> 1.44"
}

module "azure-infra" {
  source              = "../global/azure-express-route"
  resource_group_name = "terraform-acceptance-tests"
  env                 = "dev1"
}

module "google-infra" {
  source = "../global/google-cloud-interconnect"
  env    = "dev1"
}

module "pureport-infra" {
  source                   = "./pureport-infra"
  datasource_express_route = module.azure-infra.datasource_express_route
  google_compute_network   = module.google-infra.compute_network
}

