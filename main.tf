module "vpc" {
  source             = "./modules/vpc"
  cidr_block         = "10.0.0.0/16"
  public_cidr_block  = "10.0.1.0/24"
  private_cidr_block = "10.0.2.0/24"
  availability_zone  = "us-east-1a"
}

module "security_group" {
  source      = "./modules/security_group"
  name        = "web-sg"
  description = "Allow SSH"
  vpc_id      = module.vpc.vpc_id
}

module "iam" {
  source = "./modules/iam"
}

module "key_pair" {
  source = "./modules/security"
}
module "web" {
  source           = "./modules/ec2"
  ami              = data.aws_ami.ubuntu.id
  instance_type    = "t2.micro"
  instance_count   = 2
  subnet_id        = module.vpc.public_subnet_id
  sg_id            = module.security_group.sg_id
  key_name         = "tech4dev"
  role             = "web"
  enable_public_ip = true
}

module "application" {
  source           = "./modules/ec2"
  ami              = data.aws_ami.ubuntu.id
  instance_type    = "t2.micro"
  instance_count   = 2
  subnet_id        = module.vpc.public_subnet_id
  sg_id            = module.security_group.sg_id
  key_name         = "tech4dev"
  role             = "application"
  enable_public_ip = true
}

module "database" {
  source           = "./modules/ec2"
  ami              = data.aws_ami.ubuntu.id
  instance_type    = "t2.micro"
  instance_count   = 2
  subnet_id        = module.vpc.public_subnet_id
  sg_id            = module.security_group.sg_id
  key_name         = "tech4dev"
  role             = "database"
  enable_public_ip = true
}

module "bastion" {
  source           = "./modules/ec2"
  ami              = data.aws_ami.ubuntu.id
  instance_type    = "t2.micro"
  instance_count   = 1
  subnet_id        = module.vpc.public_subnet_id
  sg_id            = module.security_group.sg_id
  key_name         = "tech4dev"
  role             = "bastion"
  enable_public_ip = true
}
