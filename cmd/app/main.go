package main

import (
	// "github.com/KimNattanan/exprec-backend/internal/app"
	"github.com/KimNattanan/exprec-backend/internal/app"
	"github.com/KimNattanan/exprec-backend/pkg/aws_lambda"
	"github.com/aws/aws-lambda-go/lambda"
	fiberadapter "github.com/awslabs/aws-lambda-go-api-proxy/fiber"
)

func main() {
	var fiberLambda *fiberadapter.FiberLambda
	app.Start(&fiberLambda)
	lambda.Start(aws_lambda.Handler(fiberLambda))
}
