variable "environment" {
  description = "The environment in which the resources will be created"
  type        = string
  default     = "dev"
}

variable "app_name" {
  description = "The name of the application"
  type        = string
  default     = "myapp"
}
