package main

import (
	"context"
	dynamodb_controller "lambda/pkg/controllers/dynamodb"
	param_store_controller "lambda/pkg/controllers/parameter_store"
	s3_controller "lambda/pkg/controllers/s3"
	"log"
	"net/http"

	router "github.com/aquasecurity/lmdrouter"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var httpRouter *router.Router

func init() {
	httpRouter = router.NewRouter("/optimized", loggerMiddleware)
	httpRouter.Route("GET", "/s3", s3_controller.ListObjects)
	httpRouter.Route("GET", "/parameterstore", param_store_controller.GetParameterValue)
	httpRouter.Route("GET", "/dynamodb", dynamodb_controller.GetMovie)
}

func main() {
	lambda.Start(httpRouter.Handler)
}

func loggerMiddleware(next router.Handler) router.Handler {
	return func(ctx context.Context, req events.APIGatewayProxyRequest) (
		res events.APIGatewayProxyResponse,
		err error,
	) {
		// [LEVEL] [METHOD PATH] [CODE] EXTRA
		format := "[%s] [%s %s] [%d] %s"
		level := "INF"
		var code int
		var extra string

		res, err = next(ctx, req)
		if err != nil {
			level = "ERR"
			code = http.StatusInternalServerError
			extra = " " + err.Error()
		} else {
			code = res.StatusCode
			if code >= 400 {
				level = "ERR"
			}
		}

		log.Printf(format, level, req.HTTPMethod, req.Path, code, extra)

		return res, err
	}
}
