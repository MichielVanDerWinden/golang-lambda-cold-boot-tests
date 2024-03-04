variable "lambda_name" {
  type = string
}

variable "source_path" {
  type = string
}

variable "aws_region" {
  default = "eu-west-1"
}

variable "lambda_role_arn" {
  type = string
}

variable "lambda_environment_variables" {
  type = map(string)
}