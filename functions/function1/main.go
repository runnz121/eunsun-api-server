package main

import (
	"context"
	"eunsunapiserver/config/log"
	"eunsunapiserver/config/rds"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"go.uber.org/zap"
)

func init() {

	errRds := rds.SetUpGorm()
	if errRds != nil {
		log.Fatal("Aurora Connection error")
	}
}

func LambdaHandler(ctx context.Context, event events.APIGatewayProxyRequest) {
	log.Info("[function1]", zap.Any("request", event.RequestContext))

}

func main() {
	lambda.Start(LambdaHandler)
}
