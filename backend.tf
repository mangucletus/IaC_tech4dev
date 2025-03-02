terraform {
  required_version = "1.10.5"
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "5.86.1"
    }
  }
  # backend "s3" {
  #   bucket = "tech4dev-terraform-state"
  #   key = "main.tfstate"
  #   region = "us-east-1"
  #   dynamodb_table = "terraform-state-lock"
  # }
}
