locals {
  attributes = concat(
    [
      {
        name = var.hash_key
        type = var.range_key_type
      },
      {
        name = var.range_key
        type = var.range_key_type
      }
    ],
    var.dynamodb_attributes
  )
  from_index = length(var.range_key) > 0 ? 0 : 1

  attributes_final = slice(local.attributes, local.from_index, length(local.attributes))
}
# DynamoDB table resource

resource "aws_dynamodb_table" "default" {
  name             = "${var.environment}-${var.project}-${var.app_name}"
  billing_mode     = var.billing_mode
  hash_key         = var.hash_key
  range_key        = var.range_key
  stream_enabled   = length(var.replicas) > 0 ? true : var.enable_streams
  stream_view_type = length(var.replicas) > 0 || var.enable_streams ? var.stream_view_type : ""


  server_side_encryption {
    enabled     = var.enable_encryption
    kms_key_arn = var.server_side_encryption_kms_key_arn
  }

  point_in_time_recovery {
    enabled = var.enable_point_in_time_recovery
  }

  lifecycle {
    ignore_changes = [
      read_capacity,
      write_capacity,
      stream_enabled,
      stream_view_type,
      server_side_encryption,
      point_in_time_recovery,
    ]
  }

  dynamic "attribute" {
    for_each = local.attributes_final
    content {
      name = attribute.value.name
      type = attribute.value.type
    }
  }

  dynamic "replica" {
    for_each = var.replicas
    content {
      region_name = replica.value.region_name
    }
  }

  tags = {
    Name        = "${var.environment}-${var.project}-${var.app_name}"
    Environment = var.environment
    Project     = var.project
    Application = var.app_name
  }
}
