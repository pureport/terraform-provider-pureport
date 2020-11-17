## 1.2.0 (Unreleased)

ENHANCEMENTS:

* Add Oracle Fast Connect Data Source ([41](https://github.com/pureport/terraform-provider-pureport/pull/41) by [kevinkirkup](https://github.com/kevinkirkup))
* Automate CHANGELOG.md using paultyng/changelog-gen. ([40](https://github.com/pureport/terraform-provider-pureport/pull/40) by [kevinkirkup](https://github.com/kevinkirkup))
* Handle network transitions ([51](https://github.com/pureport/terraform-provider-pureport/pull/51) by [kevinkirkup](https://github.com/kevinkirkup))
* Update Pureport SDK v2.0.0 ([34](https://github.com/pureport/terraform-provider-pureport/pull/34) by [kevinkirkup](https://github.com/kevinkirkup))
* [PE-2083] Update Create/Delete timeouts to account for rate limiting ([46](https://github.com/pureport/terraform-provider-pureport/pull/46) by [kevinkirkup](https://github.com/kevinkirkup))
* [PE-2173] Update documentation to show ike_version as optional ([38](https://github.com/pureport/terraform-provider-pureport/pull/38) by [kevinkirkup](https://github.com/kevinkirkup))
* [PE-2210] Setup/Teardown environment for each test run ([39](https://github.com/pureport/terraform-provider-pureport/pull/39) by [kevinkirkup](https://github.com/kevinkirkup))
* Update to Terraform 0.13.x
* Update to Terraform Plugin v1.16.0

BUG FIXES:

* Set the default IKE Version to 'V2'. ([35](https://github.com/pureport/terraform-provider-pureport/pull/35) by [kevinkirkup](https://github.com/kevinkirkup))
* Update CloudServices acceptance tests based on new services. ([37](https://github.com/pureport/terraform-provider-pureport/pull/37) by [kevinkirkup](https://github.com/kevinkirkup))
* [PE-2149] Add custom validator for cloud connections for HA ([33](https://github.com/pureport/terraform-provider-pureport/pull/33) by [kevinkirkup](https://github.com/kevinkirkup))
* [PE-2257] Recreate connection on network update ([50](https://github.com/pureport/terraform-provider-pureport/pull/50) by [kevinkirkup](https://github.com/kevinkirkup))
* [PE-2316] Updating AzureRM is causing failures ([48](https://github.com/pureport/terraform-provider-pureport/pull/48) by [kevinkirkup](https://github.com/kevinkirkup))

## 1.1.9 (August 14, 2020)

ENHANCEMENTS:
  * Add support for setting `customer_asn` for AWS resources.

## 1.1.8 (March 24, 2020)

ENHANCEMENTS:
 * Update acceptance tests for new gateway naming.

## 1.1.7 (February 28, 2020)

ENHANCEMENTS:
 * Update to golang 1.14.
 * Update Acceptance tests to use the Terraform 0.12 syntax.
 * Add support for AWS Transit Peering Type.
 * Update test-infr terraform configurations to 0.12.

BUG FIXES:
 * Update the documentation for filtering data sources and resources
   to match the standard Terraform API.
 * Fix random acceptance tests failures due to changes in AWS cloud regions and services.
 * Fix TravisCI deprecation warnings.

## 1.1.6 (January 16, 2020)

ENHANCEMENTS:
 * Update to Terraform 0.12.19
 * Update to AzureRM 1.40.0
 * Update to AWS 1.60.1
 * Update to Google 2.20.1

BUG FIXES:
 * Fix issue with AWS updating Cloud Services at will
 * Increase timeouts due to changes in the backend which may result in longer expected wait times.

## 1.1.5 (November 18, 2019)

ENHANCEMENTS:
 * Update to Terraform 0.12.13
 * Update to AzureRM 1.36.1
 * Update to AWS 1.60.1
 * Updated to golangci-lint 1.21.0
 * Added Terraform Plugins SDK 1.3.0

BUG FIXES:
 * Update the billing amount to float64 to work with new REST API for 2.22.0
 * Ensure resources created during acceptance tests are unique.

## 1.1.4 (October 24, 2019)

ENHANCEMENTS:
 * Update to Terraform 0.12.10
 * Update to Golang 1.13
 * Update to AzureRM 1.35.0

BUG FIXES:
 * Added unique names for Google Cloud Interconnects and Routers for Acceptance tests
 * Remove unused Templates Provider

## 1.1.3 (September 23, 2019)

BUG FIXES:
 * Fix integration issue with ASN limits on 32bit systems

## 1.1.2 (September 12, 2019)

BUG FIXES:
 * Fixes issue with not building plugin images out of the vendor directory during CI/CD.
 * Fixes upstream build issue with the git.apache.org repositories.

## 1.1.1 (September 3, 2019)

BUG FIXES:
 * Removes the deprecated DummyConnection Resource from the website documentation.
 * Include Darwin version of the Terraform Plugin

## 1.1.0 (August 23, 2019)

NOTES:
Update to the latest Pureport golang SDK

There shouldn't be any functional differences in this release and the previous one.

## 1.0.1 (August 23, 2019)


## 1.0.0 (July 17, 2019)

NOTES:

Initial feature release of the Pureport Terraform Provider.

FEATURES:

 * Fully automated deployment of Cloud infrastructure via Terraform
 * Full documentation of all supported Data Sources and Resources
 * Automated Unit Testing and Acceptance Tests
 * Sweep target for cleaning up orphaned Pureport resources
 * Added support for custom tagging of resources
 * Added support for dynamic filtering of data sources

 * Data Sources

    pureport_accounts
    pureport_aws_connection
    pureport_azure_connection
    pureport_cloud_regions
    pureport_cloud_services
    pureport_connections
    pureport_google_cloud_connection
    pureport_locations
    pureport_networks
    pureport_site_vpn_connection

 * Resources

    resource_pureport_aws_connection
    resource_pureport_azure_connection
    resource_pureport_google_cloud_connection
    resource_pureport_network
    resource_pureport_site_vpn_connection
