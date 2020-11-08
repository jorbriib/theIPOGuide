resource "aws_security_group" "ipo-sg-db" {
  name        = "ipo-sg-db"
  description = "Segurity group for DB"
  vpc_id      = var.vpc
  ingress {
    from_port       = 3306
    to_port         = 3306
    protocol        = "tcp"
    security_groups = [var.ec2-sg]
    self            = false
  }
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
  tags = {
    Name = "ipo-sg-db"
  }
}

resource "aws_db_subnet_group" "ipo-sn-db" {
  name        = "ipo-sn-db"
  description = "Subnet for DB"
  subnet_ids  = [var.sn_priv1, var.sn_priv2]
  tags = {
    "Name" = "ipo-sn-db"

  }
}

resource "aws_db_instance" "ipo-db" {
  identifier        = "ipo-db"

  allocated_storage = 10
  storage_type      = "gp2"

  engine            = "mysql"
  engine_version    = "8.0.20"
  instance_class    = "db.t2.micro"

  name     = var.database_name
  username = var.database_user
  password = var.database_password

  db_subnet_group_name   = aws_db_subnet_group.ipo-sn-db.name
  vpc_security_group_ids = [aws_security_group.ipo-sg-db.id]
  parameter_group_name   = "default.mysql8.0"

  skip_final_snapshot       = true
  storage_encrypted = false
  multi_az = false
  auto_minor_version_upgrade = true
  publicly_accessible     = false

  backup_window = "10:00-11:30"
  backup_retention_period = 7

  tags = {
    "Name" = "ipo-db"
  }
}