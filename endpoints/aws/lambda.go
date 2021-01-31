package main

import (
	"example.com/someservice555/endpoints"
	"example.com/someservice555/internal/routing"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/gin"
)

func main() {
	engine := routing.Build()

	for _, e := range endpoints.Endpoints {
		routing.AddRoute(engine, e.Path, e.Method, e.Function)
	}

	proxy := func(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		adapter := ginadapter.New(engine)
		return adapter.Proxy(req)
	}

	lambda.Start(proxy)
}
