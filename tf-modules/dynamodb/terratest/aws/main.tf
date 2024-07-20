module "test_dynamodb" {
  source      = "../../"
  environment = var.environment
  project     = var.project
  app_name    = "bot"
}
