resource "aws_s3_bucket" "test_bucket" {
    bucket = "go-lambda-cold-boot-tests-${uuid()}"
}