variable "lambda_name" {
  type = string
}

variable "ecr_repo" {
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
