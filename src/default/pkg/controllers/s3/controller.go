package s3_controller

import (
	"context"
	aws_s3 "lambda/pkg/aws/s3"
	controller "lambda/pkg/controllers"
	"net/http"

	"github.com/aquasecurity/lmdrouter"
	"github.com/aws/aws-lambda-go/events"
)

var (
	bucket string
)

func init() {
	bucket = "michiel-golang-lambda-cold-boot-tests-tf-state"
}

/*
Lists the top 1000 objects from the bucket for the base path.
This can be enhanced by adding a Prefix in the lambda path, or using a POST and body with specific search/filter parameters.
*/
func ListObjects(ctx context.Context, req events.APIGatewayProxyRequest) (res events.APIGatewayProxyResponse, err error) {
	list, err := aws_s3.ListObjects(bucket)
	if err != nil {
		return controller.HandleError(err)
	}
	return lmdrouter.MarshalResponse(http.StatusOK, nil, list)
}
