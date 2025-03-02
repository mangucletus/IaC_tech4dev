# data "aws_ami_ids" "ubuntu" {
#   owners = ["099720109477"]

#   filter {
#     name   = "name"
#     values = ["ubuntu/images/hvm-ssd/ubuntu-*-*-amd64-server-*"]
#   }
# }

data "aws_ami" "ubuntu_ami" {
  owners = ["099720109477"]
  most_recent = true

  filter {
    name   = "name"
    values = ["ubuntu/images/hvm-ssd/ubuntu-*-*-amd64-server-*"]
  }

  filter {
    name   = "virtualization-type"
    values = ["hvm"]
  }
}

# output "ami_id" {
#   value = data.aws_ami_ids.ubuntu.ids
# }

output "ami" {
  value = data.aws_ami.ubuntu_ami.name
}
