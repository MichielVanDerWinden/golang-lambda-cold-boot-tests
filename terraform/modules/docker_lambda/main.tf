module "docker_image" {
  source = "terraform-aws-modules/lambda/aws//modules/docker-build"

  create_ecr_repo = true
  ecr_repo        = var.ecr_repo

  use_image_tag = false

  source_path = var.source_path
}

module "lambda_function" {
  source = "terraform-aws-modules/lambda/aws"

  function_name  = var.lambda_name
  create_package = false

  image_uri    = module.docker_image.image_uri
  package_type = "Image"
  lambda_role  = var.lambda_role_arn
  create_role  = false
}

output "lambda_function_arn" {
  value = module.lambda_function.lambda_function_arn
}

output "lambda_function_invoke_arn" {
  value = module.lambda_function.lambda_function_invoke_arn
}

output "lambda_function_name" {
  value = module.lambda_function.lambda_function_name
}
