resource "aws_instance" "instance" {
  count                  = var.instance_count
  ami                    = var.ami
  instance_type          = var.instance_type
  subnet_id              = var.subnet_id
  vpc_security_group_ids = [var.sg_id]
  # iam_instance_profile   = var.iam_role
  key_name = var.key_name
  associate_public_ip_address = var.enable_public_ip

  tags = {
    Name = "${var.role}-${count.index}"
    Role = var.role
    Environment = "core-banking"
  }
}

