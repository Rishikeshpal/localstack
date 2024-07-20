provider "aws" {
  region                      = "eu-west-1"
  profile                     = "default"
  s3_use_path_style           = true
  skip_credentials_validation = true
  skip_metadata_api_check     = true
  skip_requesting_account_id  = true

  endpoints {
    s3       = "http://localhost:4566"
    sts      = "http://localhost:4566"
    kms      = "http://localhost:4566"
    dynamodb = "http://localhost:4566"
    route53  = "http://localhost:4566"
    acm      = "http://localhost:4566"
  }
}

terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.1"
    }
  }
  required_version = ">= 1.0.0"
}
