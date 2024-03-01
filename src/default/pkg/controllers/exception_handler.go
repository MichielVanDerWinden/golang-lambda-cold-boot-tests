package controller

import (
	"lambda/pkg/models/exception"
	"log"

	"github.com/aquasecurity/lmdrouter"
	"github.com/aws/aws-lambda-go/events"
)

/*
Maps a given error to a properly formatted exception.
*/
func HandleError(err error) (events.APIGatewayProxyResponse, error) {
	log.Printf("An error has occurred: %s", err)
	errorView := exception.NewErrorView(err)
	return lmdrouter.MarshalResponse(errorView.ResponseCode, nil, errorView)
}
