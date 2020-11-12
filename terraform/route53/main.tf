resource "aws_route53_record" "ipo-api-pb-dns" {
  zone_id = var.dns_pb_zone
  name = "api.${var.domain}"
  type = "A"

  alias {
    evaluate_target_health = false
    name = var.lb-dns-name
    zone_id = var.lb-zone
  }
}


resource "aws_route53_record" "ipo-api-pb-dns-v6" {
  zone_id = var.dns_pb_zone
  name = "api.${var.domain}"
  type = "AAAA"

  alias {
    evaluate_target_health = false
    name = var.lb-dns-name
    zone_id = var.lb-zone
  }
}

resource "aws_route53_record" "frontend_record" {
  zone_id = var.dns_pb_zone
  name    = var.domain
  type    = "A"
  alias {
    name = var.cf-domain-name
    zone_id = var.cf-hosted-zone
    evaluate_target_health = false
  }
}

resource "aws_route53_record" "ipo-static-pb-dns" {
  zone_id = var.dns_pb_zone
  name = "static.${var.domain}"
  type = "A"

  alias {
    name = var.cf-domain-name
    zone_id = var.cf-hosted-zone
    evaluate_target_health = false
  }
}


resource "aws_route53_record" "ipo-static-pb-dns-v6" {
  zone_id = var.dns_pb_zone
  name = "static.${var.domain}"
  type = "AAAA"

  alias {
    name = var.cf-domain-name
    zone_id = var.cf-hosted-zone
    evaluate_target_health = false
  }
}

resource "aws_route53_record" "ipo-cert-validation" {
  for_each = {
    for dvo in var.acm-domain-val : dvo.domain_name => {
      name    = dvo.resource_record_name
      record  = dvo.resource_record_value
      type    = dvo.resource_record_type
    }
  }

  allow_overwrite = true
  name            = each.value.name
  records         = [each.value.record]
  ttl             = 60
  type            = each.value.type
  zone_id         = var.dns_pb_zone
}