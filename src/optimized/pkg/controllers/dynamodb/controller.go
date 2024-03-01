package dynamodb_controller

import (
	"context"
	aws_dynamodb "lambda/pkg/aws/dynamodb"
	controller "lambda/pkg/controllers"
	"net/http"

	"github.com/aquasecurity/lmdrouter"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

var ()

func init() {

}

// Create struct to hold info about new item
type Movie struct {
	Year  string `json:"year"`
	Title string `json:"title"`
	Plot  string `json:"plot"`
}

/*
Returns a movie found in the DynamoDB table based on some hardcoded inputs.
*/
func GetMovie(ctx context.Context, req events.APIGatewayProxyRequest) (res events.APIGatewayProxyResponse, err error) {
	result, err := aws_dynamodb.GetItemComposite("Movies", "Year", "2015", "Title", "The Big New Movie")
	if err != nil {
		return controller.HandleError(err)
	}
	movie := Movie{}
	err = dynamodbattribute.UnmarshalMap(result.Item, &movie)
	if err != nil {
		return controller.HandleError(err)
	}

	return lmdrouter.MarshalResponse(http.StatusOK, nil, movie)
}
