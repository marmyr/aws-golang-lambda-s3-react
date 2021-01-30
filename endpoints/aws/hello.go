package main

import (
	"example.com/someservice555/api/hello"
	"example.com/someservice555/internal/routing"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	engine := routing.Build()
	handler := routing.CreateLambdaEntrypoint(engine, hello.Path, hello.Method, hello.ProcessRequest)
	lambda.Start(handler)
}
