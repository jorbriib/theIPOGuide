provider "aws" {
  region = "us-east-1"
}

resource "aws_acm_certificate" "ipo-cert" {
  domain_name       = var.domain
  subject_alternative_names = ["*.${var.domain}"]
  validation_method = "DNS"

  tags = {
    Name = "ipo-cert"
  }

  lifecycle {
    create_before_destroy = true
  }
}

resource "aws_acm_certificate_validation" "ipo-cert-validation" {
  certificate_arn = aws_acm_certificate.ipo-cert.arn
  validation_record_fqdns = [for record in var.route53-cert-val : record.fqdn]
}
