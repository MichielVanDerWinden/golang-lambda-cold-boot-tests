package aws_common

import (
	"fmt"
	"lambda/pkg/models/exception"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

/*
Creates a connection to AWS using the specified config.
*/
func GetSession(config *aws.Config) (sess *session.Session, err error) {
	sess, err = session.NewSession(config)
	if err != nil {
		return nil, exception.InternalServerError(fmt.Sprintf("Error occurred during creation of session: %s", err.Error()))
	}
	return sess, nil
}
