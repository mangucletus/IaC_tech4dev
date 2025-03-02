output "workspace_bucket_name" {
  value = aws_s3_bucket.workspace_bucket_class.bucket
}

output "api_url" {
  value = "${aws_api_gateway_deployment.deployment.invoke_url}/hello"
}
