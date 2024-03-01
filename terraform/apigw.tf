module "api_gateway" {
  source           = "./modules/api_gw"
  api_gw_name      = "test_golang_lambda_gateway"
  api_gw_role_name = "test_golang_lambda_iam_role"
}

# Default Lambda integration
module "api_gateway_integration_default" {
  source = "./modules/api_gw_integration"

  api_gateway_resource_path    = "default"
  api_gateway_id               = module.api_gateway.api_gateway_id
  api_gateway_root_resource_id = module.api_gateway.root_resource_id
  api_gateway_execution_arn    = module.api_gateway.api_gateway_execution_arn

  lambda_function_invoke_arn = module.default_lambda.lambda_function_invoke_arn
  lambda_function_name       = module.default_lambda.lambda_function_name
}

# Optimized Lambda integration
module "api_gateway_integration_optimized" {
  source = "./modules/api_gw_integration"

  api_gateway_resource_path    = "optimized"
  api_gateway_id               = module.api_gateway.api_gateway_id
  api_gateway_root_resource_id = module.api_gateway.root_resource_id
  api_gateway_execution_arn    = module.api_gateway.api_gateway_execution_arn

  lambda_function_invoke_arn = module.optimized_lambda.lambda_function_invoke_arn
  lambda_function_name       = module.optimized_lambda.lambda_function_name
}

resource "aws_api_gateway_deployment" "default" {
  depends_on    = [module.api_gateway_integration_default, module.api_gateway_integration_optimized]
  lifecycle {
    create_before_destroy = true
  }

  rest_api_id = module.api_gateway.api_gateway_id
}

resource "aws_api_gateway_stage" "stage" {
  stage_name = "dev"
  rest_api_id = module.api_gateway.api_gateway_id
  deployment_id = aws_api_gateway_deployment.default.id
}

resource "aws_api_gateway_method_settings" "settings" {
  rest_api_id =  module.api_gateway.api_gateway_id
  stage_name  = aws_api_gateway_stage.stage.stage_name
  method_path = "*/*"
  settings {
    logging_level = "INFO"
    data_trace_enabled = true
    metrics_enabled = true 
  }
}