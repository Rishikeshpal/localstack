module "test_dynamodb" {
  source      = "../../tf-modules/dynamodb"
  environment = var.environment
  project     = var.project
  app_name    = "bot"
}
