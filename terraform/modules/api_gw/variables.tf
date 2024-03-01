variable "api_gw_name" {
  type        = string
  description = "The name of the API GW"
}

variable "api_gw_role_name" {
  type        = string
  description = "The name of the API GW IAM Role"
}

variable "aws_region" {
  default = "eu-west-1"
}