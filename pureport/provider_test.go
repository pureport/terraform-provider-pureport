package pureport

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/terraform-providers/terraform-provider-aws/aws"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm"
	"github.com/terraform-providers/terraform-provider-google/google"
)

var (
	testAccProviders map[string]terraform.ResourceProvider
	testAccProvider  *schema.Provider

	testEnvironmentName     string = "Production"
	testOraclePrimaryOCID   string = "ocid1.virtualcircuit.oc1.iad.aaaaaaaa0000000000000000000000000000000000000000000000000001"
	testOracleSecondaryOCID string = "ocid1.virtualcircuit.oc1.iad.aaaaaaaa0000000000000000000000000000000000000000000000000002"
)

func init() {
	testAccProvider = Provider().(*schema.Provider)

	testAccProviders = map[string]terraform.ResourceProvider{
		"pureport": testAccProvider,
		"google":   google.Provider(),
		"aws":      aws.Provider(),
		"azurerm":  azurerm.Provider(),
	}

	// Environment Variables for the Test Environment
	if env := os.Getenv("PUREPORT_ACC_TEST_ENVIRONMENT"); env != "" {
		testEnvironmentName = env
	}

	if ocid := os.Getenv("TF_VAR_oracle_primary_ocid"); ocid != "" {
		testOraclePrimaryOCID = ocid
	}

	if ocid := os.Getenv("TF_VAR_oracle_secondary_ocid"); ocid != "" {
		testOracleSecondaryOCID = ocid
	}
}

func testAccPreCheck(t *testing.T) {

	pureportEnvVars := []string{
		"PUREPORT_API_KEY",
		"PUREPORT_API_SECRET",
		"PUREPORT_ENDPOINT",
	}

	googleEnvVars := []string{
		"GOOGLE_CREDENTIALS",
		"GOOGLE_PROJECT",
		"GOOGLE_REGION",
	}

	amazonEnvVars := []string{
		"AWS_DEFAULT_REGION",
		"AWS_ACCESS_KEY_ID",
		"AWS_SECRET_ACCESS_KEY",
	}

	azureEnvVars := []string{
		"ARM_CLIENT_ID",
		"ARM_CLIENT_SECRET",
		"ARM_SUBSCRIPTION_ID",
		"ARM_TENANT_ID",
		"ARG_USE_MSI",
	}

	// Pureport Provider Configuration
	for _, e := range pureportEnvVars {
		if v := os.Getenv(e); v == "" {
			t.Fatalf("%s must be specified for acceptance tests", e)
		}
	}

	// Google Cloud Provider Configuration
	for _, e := range googleEnvVars {
		if v := os.Getenv(e); v == "" {
			t.Fatalf("%s must be specified for acceptance tests", e)
		}
	}

	// AWS Cloud Provider Configuration
	for _, e := range amazonEnvVars {
		if v := os.Getenv(e); v == "" {
			t.Fatalf("%s must be specified for acceptance tests", e)
		}
	}

	// Azure Cloud Provider Configuration
	for _, e := range azureEnvVars {
		if v := os.Getenv(e); v == "" {
			t.Fatalf("%s must be specified for acceptance tests", e)
		}
	}

	if err := testAccProvider.Configure(terraform.NewResourceConfigRaw(nil)); err != nil {
		t.Fatal(err)
	}
}
