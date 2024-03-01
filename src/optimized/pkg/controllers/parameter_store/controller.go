package param_store_controller

import (
	"context"
	aws_parameter_store "lambda/pkg/aws/parameter_store"
	controller "lambda/pkg/controllers"
	"net/http"

	"github.com/aquasecurity/lmdrouter"
	"github.com/aws/aws-lambda-go/events"
)

type getParameterValueOutput struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

/*
Returns the parameter found for a hardcoded set of inputs.
*/
func GetParameterValue(ctx context.Context, req events.APIGatewayProxyRequest) (res events.APIGatewayProxyResponse, err error) {
	parameterName := "/test/param"
	parameterValue, err := aws_parameter_store.GetParameterValue(parameterName)
	if err != nil {
		return controller.HandleError(err)
	}
	return lmdrouter.MarshalResponse(http.StatusOK, nil, getParameterValueOutput{
		Name:  parameterName,
		Value: *parameterValue,
	})
}
