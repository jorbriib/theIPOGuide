output "ec2" {
  value = aws_instance.ipo-ec2.id
}
output "ec2-sg" {
  value = aws_security_group.ipo-sg-priv-ec2.id
}