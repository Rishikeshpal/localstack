resource "aws_route53_zone" "example" {
  name = "seturgoal.com"
}

resource "aws_route53_record" "example" {
  allow_overwrite = true
  name            = aws_route53_zone.example.name
  ttl             = 172800
  type            = "NS"
  zone_id         = aws_route53_zone.example.zone_id

  records = [
    aws_route53_zone.example.name_servers[0],
    aws_route53_zone.example.name_servers[1],
    aws_route53_zone.example.name_servers[2],
    aws_route53_zone.example.name_servers[3],
  ]
}

resource "aws_acm_certificate" "example" {
  domain_name       = "seturgoal.com"
  validation_method = "DNS"

  lifecycle {
    create_before_destroy = true
  }
}

resource "aws_kms_key" "example" {
  deletion_window_in_days = 10
}

resource "aws_kms_alias" "example" {
  name          = "alias/example"
  target_key_id = aws_kms_key.example.key_id
}

output "zone_id" {
  value = aws_route53_zone.example.zone_id
}

output "zone_name" {
  value = aws_route53_zone.example.name
}

output "name_servers" {
  value = aws_route53_zone.example.name_servers
}

output "certificate_arn" {
  value = aws_acm_certificate.example.arn
}

output "kms_key_id" {
  value = aws_kms_key.example.key_id
}

output "kms_alias_name" {
  value = aws_kms_alias.example.name
}
