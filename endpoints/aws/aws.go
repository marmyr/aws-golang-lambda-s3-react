package main

import (
	"example.com/someservice555/api/hello"
	"example.com/someservice555/internal/routing"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/gin"
)

func main() {
	engine := routing.Build()

	routing.AddRoute(engine, hello.Path, hello.Method, hello.ProcessRequest)

	proxy := func(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		adapter := ginadapter.New(engine)
		return adapter.Proxy(req)
	}

	lambda.Start(proxy)
}
