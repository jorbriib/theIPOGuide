output "acm-domain-val" {
  value = aws_acm_certificate.ipo-cert.domain_validation_options
}

output "ssl-cert-arn" {
  value = aws_acm_certificate.ipo-cert.arn
}