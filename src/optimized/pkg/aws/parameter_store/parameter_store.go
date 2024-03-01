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
}

/*
Returns a ssmSvc connection, created with a new session.
Reuses the previously created ssmSvc if possible.
*/
func getService() (*ssm.SSM, error) {
	if ssmSvc == nil {
		time.Sleep(time.Second)
		sess, err := aws_common.GetSession(config)
		if err != nil {
			return nil, err
		}
		ssmSvc = ssm.New(sess, config)
	}
	return ssmSvc, nil
}

/*
Returns the value of the parameter found by the given parameterName.
*/
func GetParameterValue(parameterName string) (parameterValue *string, err error) {
	ssmSvc, err = getService()
	if err != nil {
		return
	}
	parameter, err := ssmSvc.GetParameter(&ssm.GetParameterInput{
		Name:           aws.String(parameterName),
		WithDecryption: aws.Bool(true),
	})
	if err != nil {
		return
	}
	parameterValue = parameter.Parameter.Value
	return
}
