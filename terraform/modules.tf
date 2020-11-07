module "network" {
  source = "./network"
  region = var.region
  domain = var.domain

}
module "ec2" {
  source = "./ec2"
  region = var.region
  vpc = module.network.vpc
  sn_pub1 = module.network.sn_pub1
  key_name = var.key_name
  lb-security-group = module.alb.lb-security-group

}
module "alb" {
  source = "./alb"
  region = var.region
  vpc = module.network.vpc
  sn_pub1 = module.network.sn_pub1
  sn_pub2 = module.network.sn_pub2
  ec2 = module.ec2.ec2
  domain = var.domain
}

module "db" {
  source = "./db"
  region = var.region
  vpc = module.network.vpc
  ec2-sg = module.ec2.ec2-sg
  sn_priv1 = module.network.sn_priv1
  sn_priv2 = module.network.sn_priv2

  database_name = var.database_name
  database_user = var.database_user
  database_password = var.database_password
}

module "acm" {
  source = "./acm"
  domain = var.domain
  route53-cert-val = module.route53.route53-cert-val
}

module "s3" {
  source = "./s3"
  domain = var.domain
}

module "cloudfront" {
  source = "./cloudfront"

  domain = var.domain
  s3-website-endpoint = module.s3.s3-website-endpoint
  ssl-cert-arn = module.acm.ssl-cert-arn

}

module "route53" {
  source = "./route53"
  region = var.region
  domain = var.domain
  vpc = module.network.vpc
  dns_pb_zone = var.dns_pb_zone

  lb-dns-name = module.alb.lb-dns-name
  lb-zone = module.alb.lb-zone

  cf-domain-name = module.cloudfront.cf-domain-name
  cf-hosted-zone = module.cloudfront.cf-hosted-zone


  acm-domain-val = module.acm.acm-domain-val
}