resource "aws_dynamodb_table_item" "big_new_movie" {
  table_name = aws_dynamodb_table.movies.name
  hash_key   = aws_dynamodb_table.movies.hash_key
  range_key  = aws_dynamodb_table.movies.range_key

  item = <<ITEM
{
  "Year": {"S": "2015"},
  "Title": {"S": "The Big New Movie"},
  "Plot": {"S": "Nothing happens at all."}
}
ITEM
}

resource "aws_dynamodb_table" "movies" {
  name         = "Movies"
  hash_key     = "Year"
  range_key    = "Title"
  billing_mode = "PAY_PER_REQUEST"

  attribute {
    name = "Year"
    type = "S"
  }

  attribute {
    name = "Title"
    type = "S"
  }
}

resource "aws_ssm_parameter" "test_param" {
  name  = "/test/param"
  type  = "String"
  value = "my_test_param_value"
}

resource "aws_iam_role" "lambda_role" {
  name               = "test_golang_lambda_role"
  assume_role_policy = data.aws_iam_policy_document.lambda_assume_role.json
}

resource "aws_iam_policy" "lambda_iam_policy" {
  name   = "test_golang_lambda_iam_policy"
  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": [
        "dynamodb:*"
      ],
      "Resource": [
        "${aws_dynamodb_table.movies.arn}"
      ],
      "Effect": "Allow"
    },
    {
      "Action": [
        "s3:*"
      ],
      "Resource": [
        "${aws_s3_bucket.terraform_state.arn}/*",
        "${aws_s3_bucket.terraform_state.arn}"
      ],
      "Effect": "Allow"
    },
    {
      "Action": [
        "ssm:*"
      ],
      "Resource": [
        "${aws_ssm_parameter.test_param.arn}"
      ],
      "Effect": "Allow"
    },
    {
      "Action": [
        "cloudwatch:*"
      ],
      "Resource": "*",
      "Effect": "Allow"
    }
  ]
}
  EOF
}

resource "aws_iam_role_policy_attachment" "attach_iam_policy_to_iam_role" {
  role       = aws_iam_role.lambda_role.name
  policy_arn = aws_iam_policy.lambda_iam_policy.arn
}

module "default_lambda" {
  source = "./modules/docker_lambda"

  ecr_repo    = "default_golang_lambda"
  lambda_name = "default_golang_lambda_function"
  source_path = "../src/default"
  lambda_role_arn = aws_iam_role.lambda_role.arn
}

module "optimized_lambda" {
  source = "./modules/docker_lambda"

  ecr_repo    = "optimized_golang_lambda"
  lambda_name = "optimized_golang_lambda_function"
  source_path = "../src/optimized"
  lambda_role_arn = aws_iam_role.lambda_role.arn
}