package aws_parameter_store

import (
	aws_common "lambda/pkg/aws/_common"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ssm"
)

var (
	config *aws.Config
	ssmSvc *ssm.SSM
)

func init() {
	config = aws_common.GetAwsConfiguration()
	time.Sleep(time.Second)
	sess, _ := aws_common.GetSession(config)
	ssmSvc = ssm.New(sess, config)
}

/*
Returns the value of the parameter found by the given parameterName.
*/
func GetParameterValue(parameterName string) (parameterValue string, err error) {
	parameter, err := ssmSvc.GetParameter(&ssm.GetParameterInput{
		Name:           aws.String(parameterName),
		WithDecryption: aws.Bool(true),
	})
	if err != nil {
		return
	}
	parameterValue = *parameter.Parameter.Value
	return
}
