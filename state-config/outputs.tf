output "s3_bucket_name" {
  description = "S3 Bucket Name for Terraform State"
  value       = aws_s3_bucket.terraform_state.bucket
}

output "dynamodb_table_name" {
  description = "DynamoDB Table Name for Terraform State Locking"
  value       = aws_dynamodb_table.terraform_locks.name
}

output "iam_user_arn" {
  description = "IAM User ARN for Terraform"
  value       = aws_iam_user.terraform_user.arn
}

output "iam_policy_arn" {
  description = "IAM Policy ARN for Terraform Backend Access"
  value       = aws_iam_policy.terraform_backend_policy.arn
}
