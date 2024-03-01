package aws_dynamodb

import (
	aws_common "lambda/pkg/aws/_common"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var (
	config      *aws.Config
	dynamodbSvc *dynamodb.DynamoDB
)

func init() {
	config = aws_common.GetAwsConfiguration()
	time.Sleep(time.Second)
	sess, _ := aws_common.GetSession(config)
	dynamodbSvc = dynamodb.New(sess)
}

/*
Returns an item found by the given parition key value.
*/
func GetItem(tableName string, partitionName string, partitionValue string) (result *dynamodb.GetItemOutput, err error) {
	result, err = dynamodbSvc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			partitionName: {
				S: aws.String(partitionValue),
			},
		},
	})
	return
}

/*
Returns an item found by the given composite key values.
*/
func GetItemComposite(tableName string, partitionName string, partitionValue string, sortName string, sortValue string) (result *dynamodb.GetItemOutput, err error) {
	result, err = dynamodbSvc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			partitionName: {
				S: aws.String(partitionValue),
			},
			sortName: {
				S: aws.String(sortValue),
			},
		},
	})
	return
}
