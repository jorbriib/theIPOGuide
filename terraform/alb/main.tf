resource "aws_security_group" "ipo-alb-sg" {
  name   = "ipo-alb-sg"
  vpc_id = var.vpc

  ingress {
    protocol          = "tcp"
    from_port         = 80
    to_port           = 80
    cidr_blocks       = ["0.0.0.0/0"]
  }

  ingress {
    protocol          = "tcp"
    from_port         = 443
    to_port           = 443
    cidr_blocks       = ["0.0.0.0/0"]
  }

  egress {
    protocol = "-1"
    from_port = 0
    to_port = 0
    cidr_blocks       = ["0.0.0.0/0"]
  }
  tags = {
    Name = "ipo-alb-sg"
  }
}

resource "aws_lb" "ipo-lb" {
  name = "ipo-lb"
  internal = false
  load_balancer_type = "application"
  security_groups = [aws_security_group.ipo-alb-sg.id]
  subnets = [var.sn_pub1, var.sn_pub2]

  enable_deletion_protection = false
}

resource "aws_lb_listener" "ipo-lb-listener-80" {
  load_balancer_arn = aws_lb.ipo-lb.arn
  port = "80"
  protocol = "HTTP"

  default_action {
    type = "redirect"

    redirect {
      port = "443"
      protocol = "HTTPS"
      status_code = "HTTP_301"
    }
  }
}

resource "aws_lb_target_group" "ipo-lb-target-group-80" {
  name = "ipo-lb-target-group-80"
  port = "80"
  protocol = "HTTP"
  vpc_id = var.vpc

  health_check {
    enabled = true
    path = "/health-check"
    matcher = "200,204"
  }
}

resource "aws_lb_target_group_attachment" "ipo-lb-target-group-attach-ec2" {
  target_group_arn = aws_lb_target_group.ipo-lb-target-group-80.arn
  target_id = var.ec2
  port = 80
}

data "aws_acm_certificate" "ipo-cert-region" {
  domain   = var.domain
  statuses = ["ISSUED"]
}

resource "aws_lb_listener" "ipo-lb-listener-443" {
  load_balancer_arn = aws_lb.ipo-lb.arn
  port = "443"
  protocol = "HTTPS"
  certificate_arn = data.aws_acm_certificate.ipo-cert-region.arn
  ssl_policy = "ELBSecurityPolicy-2016-08"
  default_action {
    type = "forward"
    target_group_arn = aws_lb_target_group.ipo-lb-target-group-80.arn
  }
}
