variable "aws_region" {
  description = "The AWS region"
  default     = "us-east-1"
}

variable "bucket_name" {
  description = "The name of the S3 bucket for the Terraform state"
  default     = "tech4dev-terraform-state"
}

variable "dynamodb_table" {
  description = "The name of the DynamoDB table for state locking"
  default     = "terraform-state-lock"
}

variable "iam_user_name" {
  description = "The name of the IAM user for Terraform"
  default     = "kk_labs_user_677738"
}
