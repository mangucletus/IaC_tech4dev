variable "ami" {}
variable "instance_type" {}
variable "instance_count" {}
variable "subnet_id" {}
variable "sg_id" {}
# variable "iam_role" {}
variable "key_name" {}
variable "role" {}
variable "enable_public_ip" {
    default = false
}

