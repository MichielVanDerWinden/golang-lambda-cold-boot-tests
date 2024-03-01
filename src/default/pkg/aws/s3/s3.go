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
	time.Sleep(time.Second)
	sess, _ := aws_common.GetSession(config)
	s3Svc = s3.New(sess, config)
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
	return objects, nil
}
