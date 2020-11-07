resource "aws_vpc" "ipo-vcp" {
  cidr_block           = "10.0.0.0/16"
  enable_dns_hostnames = false
  enable_dns_support   = true
  instance_tenancy     = "default"
  tags = {
    Name = "ipo-vpc"
  }
}

data "aws_availability_zones" "available" {}


resource "aws_subnet" "ipo-sn-public-1" {
  vpc_id                  = aws_vpc.ipo-vcp.id
  cidr_block              = "10.0.1.0/24"
  availability_zone       = data.aws_availability_zones.available.names[0]
  map_public_ip_on_launch = true
  tags = {
    Name = "ipo-sn-public-1"
  }
}

resource "aws_subnet" "ipo-sn-public-2" {
  vpc_id                  = aws_vpc.ipo-vcp.id
  cidr_block              = "10.0.2.0/24"
  availability_zone       = data.aws_availability_zones.available.names[1]
  map_public_ip_on_launch = true
  tags = {
    Name = "ipo-sn-public-2"
  }
}


resource "aws_subnet" "ipo-sn-private-1" {
  vpc_id                  = aws_vpc.ipo-vcp.id
  cidr_block              = "10.0.4.0/24"
  availability_zone       = data.aws_availability_zones.available.names[0]
  map_public_ip_on_launch = true
  tags = {
    Name = "ipo-sn-private-1"
  }
}


resource "aws_subnet" "ipo-sn-private-2" {
  vpc_id                  = aws_vpc.ipo-vcp.id
  cidr_block              = "10.0.3.0/24"
  availability_zone       = data.aws_availability_zones.available.names[1]
  map_public_ip_on_launch = true
  tags = {
    Name = "ipo-sn-private-2"
  }
}
resource "aws_internet_gateway" "ipo-igw" {
  vpc_id = aws_vpc.ipo-vcp.id
  tags = {
    Name = "ipo-igw"
  }
}

resource "aws_route_table" "ipo-rt" {
  vpc_id = aws_vpc.ipo-vcp.id
  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.ipo-igw.id
  }
  tags = {
    Name = "ipo-rt"
  }
}

resource "aws_route_table_association" "ipo-rta-public-1" {
  subnet_id      = aws_subnet.ipo-sn-public-1.id
  route_table_id = aws_route_table.ipo-rt.id
}
resource "aws_route_table_association" "ipo-rta-public-2" {
  subnet_id      = aws_subnet.ipo-sn-public-2.id
  route_table_id = aws_route_table.ipo-rt.id
}

/*
resource "aws_acm_certificate" "ipo-cert" {
  domain_name       = var.domain
  validation_method = "DNS"
  lifecycle {
    create_before_destroy = true
  }
}

resource "aws_acm_certificate_validation" "ipo-cert-validation" {
  certificate_arn         = aws_acm_certificate.ipo-cert.arn
  validation_record_fqdns = [ aws_route53_record.ipo-cert-validation.fqdn ]
}

resource "aws_route53_record" "ipo-cert-validation" {
  zone_id = var.dns_pb_zone
  name    = aws_acm_certificate.ipo-cert.domain_validation_options.0.resource_record_name
  type    = aws_acm_certificate.ipo-cert.domain_validation_options.0.resource_record_type
  ttl     = 60

  records = [ aws_acm_certificate.ipo-cert.domain_validation_options.0.resource_record_value ]
}*/