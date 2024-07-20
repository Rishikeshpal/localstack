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
		TerraformDir: "./aws/",

		// Variables to pass to our Terraform code using -var-file options
		VarFiles: []string{"varfile.tfvars"},

		// Disable colors in Terraform commands so its easier to parse stdout/stderr
		NoColor: true,

		// Test targets
		Targets: []string{target},
	})
}

func TestTerraformDynamodb(t *testing.T) {

	terraformOptions := configureTerraformOptions(t, "module.test_dynamodb")

	// Using "defer" runs the command at the end of the test, whether the test succeeds or fails.
	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terraformOptions)

	// This will run `terraform init` and `terraform apply` and fail the test if there are any errors
	terraform.InitAndApply(t, terraformOptions)

	// Run `terraform output` to get the values of output variables
	tableID := terraform.Output(t, terraformOptions, "dynamodb_table_id")
	tableARN := terraform.Output(t, terraformOptions, "dynamodb_table_arn")
	tableNAME := terraform.Output(t, terraformOptions, "dynamodb_table_name")

	// Output testing
	assert.Contains(t, tableID, "aws-modules-testing")
	assert.Contains(t, tableARN, "arn:aws:dynamodb:eu-west-1:")
	assert.Contains(t, tableNAME, "aws-modules-testing")
}
