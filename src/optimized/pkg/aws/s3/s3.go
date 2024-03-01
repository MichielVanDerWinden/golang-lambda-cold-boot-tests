package aws_s3

import (
	aws_common "lambda/pkg/aws/_common"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

var (
	config *aws.Config
	s3Svc  *s3.S3
)

func init() {
	config = aws_common.GetAwsConfiguration()
}

/*
Returns an s3Svc connection, created with a new session.
Reuses the previously created s3Svc if possible.
*/
func getService() (*s3.S3, error) {
	if s3Svc == nil {
		time.Sleep(time.Second)
		sess, err := aws_common.GetSession(config)
		if err != nil {
			return nil, err
		}
		s3Svc = s3.New(sess, config)
	}
	return s3Svc, nil
}

type S3Object struct {
	Key         string
	LastUpdated time.Time
}

/*
List the top 1000 objects for the top level prefix in the given bucket.
*/
func ListObjects(bucket string) (objects []S3Object, err error) {
	return ListObjectWithPrefix(bucket, "")
}

/*
List the top 1000 objects for the given prefix in the given bucket.
*/
func ListObjectWithPrefix(bucket string, prefix string) (objects []S3Object, err error) {
	s3Svc, err = getService()
	if err != nil {
		return
	}
	output, err := s3Svc.ListObjects(&s3.ListObjectsInput{
		Bucket: aws.String(bucket),
		Prefix: aws.String(prefix),
	})
	if err != nil {
		return
	}
	for _, i := range output.Contents {
		objects = append(objects, S3Object{
			Key:         *i.Key,
			LastUpdated: *i.LastModified,
		})
	}
	return
}
