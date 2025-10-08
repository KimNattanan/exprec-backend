package aws_lambda

import (
	"os"

	"github.com/aws/aws-lambda-go/events"
	fiberadapter "github.com/awslabs/aws-lambda-go-api-proxy/fiber"
)

type Response struct {
	Message string `json:"message"`
}

func Handler(fiberLambda *fiberadapter.FiberLambda) interface{} {
	return func(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		resp, err := fiberLambda.Proxy(request)
		origin := os.Getenv("FRONTEND_URL")

		if resp.Headers == nil {
			resp.Headers = map[string]string{}
		}

		// Ensure CORS headers are present in ALL responses
		resp.Headers["Access-Control-Allow-Origin"] = origin
		resp.Headers["Access-Control-Allow-Credentials"] = "true"
		resp.Headers["Vary"] = "Origin"

		// Handle preflight (OPTIONS) manually
		if request.HTTPMethod == "OPTIONS" {
			resp.StatusCode = 200
			resp.Body = ""
			resp.Headers["Access-Control-Allow-Methods"] = "GET,POST,PUT,PATCH,DELETE,OPTIONS"
			resp.Headers["Access-Control-Allow-Headers"] = "Origin, Content-Type, Accept, Authorization"
		}

		return resp, err
	}
}
