package aws_common

import "github.com/aws/aws-sdk-go/aws"

/*
Returns an AWS Configuration with specific settings for our usecase.
*/
func GetAwsConfiguration() *aws.Config {
	return aws.NewConfig().WithRegion("eu-west-1")
}
