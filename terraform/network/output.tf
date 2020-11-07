output "vpc" {
  value = aws_vpc.ipo-vcp.id
}
output "sn_pub1" {
  value = aws_subnet.ipo-sn-public-1.id
}
output "sn_pub2" {
  value = aws_subnet.ipo-sn-public-2.id
}
output "sn_priv1" {
  value = aws_subnet.ipo-sn-private-1.id
}
output "sn_priv2" {
  value = aws_subnet.ipo-sn-private-2.id
}