# module "docker_image" {
#   source = "terraform-aws-modules/lambda/aws//modules/docker-build"

#   create_ecr_repo = true
#   ecr_repo        = var.ecr_repo

#   use_image_tag = false

#   source_path = var.source_path
# }

// build the binary for the lambda function in a specified path
resource "null_resource" "function_binary" {
  provisioner "local-exec" {
    working_dir = var.source_path
    command = "GOOS=linux CGO_ENABLED=0 go build -ldflags=\"-w -s\" -tags lambda.norpc -o bootstrap"
  }
}

module "lambda_function" {
  depends_on = [ null_resource.function_binary ]
  source = "terraform-aws-modules/lambda/aws"

  function_name  = var.lambda_name
  create_package = true

  lambda_role = var.lambda_role_arn
  create_role = false

  source_path = var.source_path
  handler = "bootstrap"
  runtime = "provided.al2"

  environment_variables = var.lambda_environment_variables
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
