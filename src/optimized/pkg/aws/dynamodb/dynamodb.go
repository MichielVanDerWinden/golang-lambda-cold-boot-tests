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
}

/*
Returns a dynamodbSvc connection, created with a new session.
Reuses the previously created dynamodbSvc if possible.
*/
func getService() (*dynamodb.DynamoDB, error) {
	if dynamodbSvc == nil {
		time.Sleep(time.Second)
		sess, err := aws_common.GetSession(config)
		if err != nil {
			return nil, err
		}
		dynamodbSvc = dynamodb.New(sess)
	}
	return dynamodbSvc, nil
}

/*
Returns an item found by the given parition key value.
*/
func GetItem(tableName string, partitionName string, partitionValue string) (result *dynamodb.GetItemOutput, err error) {
	dynamodbSvc, err = getService()
	if err != nil {
		return nil, err
	}
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
	dynamodbSvc, err = getService()
	if err != nil {
		return
	}
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
