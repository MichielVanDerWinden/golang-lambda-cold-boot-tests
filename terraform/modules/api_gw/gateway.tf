resource "aws_api_gateway_rest_api" "gateway" {
  name = var.api_gw_name
  endpoint_configuration {
    types = ["REGIONAL"]
  }
}

output "api_gateway_id" {
  value = aws_api_gateway_rest_api.gateway.id
}

output "root_resource_id" {
  value = aws_api_gateway_rest_api.gateway.root_resource_id
}

output "api_gateway_execution_arn" {
  value = aws_api_gateway_rest_api.gateway.execution_arn
}