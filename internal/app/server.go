package app

import (
	"log"
	"os"

	"github.com/KimNattanan/exprec-backend/pkg/aws_lambda"
	"github.com/aws/aws-lambda-go/lambda"
	fiberadapter "github.com/awslabs/aws-lambda-go-api-proxy/fiber"
)

func Start() {
	db, err := setupDependencies("development")
	if err != nil {
		log.Fatalf("failed to setup dependencies: %v", err)
	}
	app := setupRestServer(db)

	if os.Getenv("ENV") == "production" {
		fiberLambda := fiberadapter.New(app)
		lambda.Start(aws_lambda.Handler(fiberLambda))
	} else {
		app.Listen(":8000")
	}
}
