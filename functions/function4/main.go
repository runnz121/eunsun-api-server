package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func LambdaHandler(ctx context.Context, event events.APIGatewayProxyRequest) {

}

func main() {
	lambda.Start(LambdaHandler)
}
