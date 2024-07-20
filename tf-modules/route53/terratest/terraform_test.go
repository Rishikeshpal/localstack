package test

// github.com/gruntwork-io/terratest v0.40.6
import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

const LocalstackEndpoint = "http://localhost:4566"
const REGION = "eu-west-1"

var s3session *s3.S3

func init() {
	s3session = s3.New(session.Must(session.NewSession(&aws.Config{
		Region:           aws.String(REGION),
		Endpoint:         aws.String(LocalstackEndpoint),
		S3ForcePathStyle: aws.Bool(true),
	})))
}

func configureTerraformOptions(t *testing.T, target string) *terraform.Options {
	return terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: "../",

		// Variables to pass to our Terraform code using -var-file options
		VarFiles: []string{"../tfvar.tfvars"},

		// Disable colors in Terraform commands so its easier to parse stdout/stderr
		NoColor: true,

		// Test targets
		Targets: []string{target},
	})
}
func TestTerraformRoute53(t *testing.T) {
	t.Parallel()

	// Define the Terraform options
	terraformOptions := &terraform.Options{
		// Path to the Terraform code being tested
		TerraformDir: "../",

		// Variables to pass to the Terraform code
		VarFiles: []string{"tfvar.tfvars"},
		// Add any variables required by your Terraform code here
	}

	// Clean up resources with "terraform destroy" at the end of the test
	defer terraform.Destroy(t, terraformOptions)

	// Run "terraform init" and "terraform apply". Fail the test if there are any errors.
	terraform.InitAndApply(t, terraformOptions)

	// Run `terraform output` to get the values of output variables
	//zoneID := terraform.Output(t, terraformOptions, "zone_id")
	zoneName := terraform.Output(t, terraformOptions, "zone_name")

	// Validate the Route 53 record exists
	assert.Contains(t, "seturgoal.com", zoneName)

}
