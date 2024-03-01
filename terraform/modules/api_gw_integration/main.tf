resource "aws_lambda_permission" "lambda_invoke_permission" {
  statement_id  = "AllowAPIGatewayInvoke"
  action        = "lambda:InvokeFunction"
  function_name = var.lambda_function_name
  principal     = "apigateway.amazonaws.com"
  source_arn    = "${var.api_gateway_execution_arn}/*/*"
}

resource "aws_api_gateway_resource" "resource" {
  path_part   = var.api_gateway_resource_path
  parent_id   = var.api_gateway_root_resource_id
  rest_api_id = var.api_gateway_id
}

resource "aws_api_gateway_resource" "proxy_resource" {
  path_part   = "{proxy+}"
  parent_id   = aws_api_gateway_resource.resource.id
  rest_api_id = var.api_gateway_id
}

resource "aws_api_gateway_method" "proxy_method" {
  rest_api_id      = var.api_gateway_id
  resource_id      = aws_api_gateway_resource.proxy_resource.id
  http_method      = "ANY"
  authorization    = "NONE"
  api_key_required = false
}

resource "aws_api_gateway_integration" "proxy_integration" {
  rest_api_id             = var.api_gateway_id
  resource_id             = aws_api_gateway_resource.proxy_resource.id
  http_method             = aws_api_gateway_method.proxy_method.http_method
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = var.lambda_function_invoke_arn
}