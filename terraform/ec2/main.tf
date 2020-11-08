resource "aws_security_group" "ipo-sg-priv-ec2" {
  name        = "ipo-sg-ec2"
  description = "Segurity group for EC2"
  vpc_id      = var.vpc
  ingress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
  ingress {
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    security_groups = [var.lb-security-group]
  }
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
  tags = {
    Name = "ipo-sg-ec2"
  }
}


resource "aws_instance" "ipo-ec2" {
  ami           = "ami-00a205cb8e06c3c4e"
  instance_type = "t2.micro"
  key_name      = var.key_name
  user_data     = file(format("%s/%s", path.module, "user_data.sh"))
  subnet_id     = var.sn_pub1
  associate_public_ip_address = true
  vpc_security_group_ids = [aws_security_group.ipo-sg-priv-ec2.id]
  iam_instance_profile   = aws_iam_instance_profile.ipo-profile.name
  tags = {
    "Name" = "ipo-ec2"
  }
}

resource "aws_iam_role" "ipo-role" {
  name = "ipo-role"

  assume_role_policy = <<EOF
{
     "Version": "2012-10-17",
     "Statement": [
       {
         "Action": "sts:AssumeRole",
         "Principal": {
         "Service": "ec2.amazonaws.com"
       },
         "Effect": "Allow",
         "Sid": ""
       }
     ]
}
EOF
}

resource "aws_iam_role_policy_attachment" "ipo-policy-ecr-read-only" {
  role       = aws_iam_role.ipo-role.name
  policy_arn = "arn:aws:iam::aws:policy/AmazonEC2ContainerRegistryReadOnly"
}

resource "aws_iam_instance_profile" "ipo-profile" {
  name  = "ipo-profile"
  role = aws_iam_role.ipo-role.name
}

resource "aws_ecr_repository" "ipo-ecr-backend-repo" {
  name                 = "ipo-ecr-backend-repo"
  image_tag_mutability = "MUTABLE"

  image_scanning_configuration {
    scan_on_push = false
  }
}

resource "aws_ecr_repository" "ipo-ecr-migrations-repo" {
  name                 = "ipo-ecr-migrations-repo"
  image_tag_mutability = "MUTABLE"

  image_scanning_configuration {
    scan_on_push = false
  }
}
