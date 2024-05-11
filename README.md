# localstack
Test your terraform code using terratest and localstack

LocalStack is a cloud service emulator that runs in a single container on your laptop or in your CI environment. With LocalStack, you can run your AWS applications or Lambdas entirely on your local machine without connecting to a remote cloud provider!

Whether you are testing complex CDK applications or Terraform configurations, or just beginning to learn about AWS services, LocalStack helps speed up and simplify your testing and development workflow.

LocalStack supports a growing number of AWS services, like AWS Lambda, S3, DynamoDB, Kinesis, SQS, SNS, and more! LocalStack Pro supports additional APIs and advanced features to make your cloud development experience a breeze!
ou can find a comprehensive list of supported APIs on our Feature Coverage page.
https://docs.localstack.cloud/user-guide/aws/feature-coverage/

# Install localstack
On macos:
```
brew  install localstack
```
# How to launch localstack
1) cli 'SERVICES=s3,sts localstack start'
2) docker-compose `cd docker && docker-compose up`  or `docker-compose -f docker/docker-compose.yml up`

# Verify the status of localstack
run the command `curl --silent --output /dev/null --write-out "%{http_code}" http://localhost:4566` and ths status should 200, that means the localstack is running

# Run the test
```
cd test/terratest
go mod tidy
go test -v
```
