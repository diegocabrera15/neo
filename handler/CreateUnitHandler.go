package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/go/neo/service"
	"github.com/go/neo/types"
)

// ResponseUnit is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type ResponseUnit events.APIGatewayProxyResponse

// HandlerCreateUnit is our lambda handler invoked by the `lambda.Start` function call
func HandlerCreateUnit(ctx context.Context, event events.APIGatewayProxyRequest) (ResponseUnit, error) {
	var buf bytes.Buffer

	service.CreateUnit(event)

	body, err := json.Marshal(map[string]interface{}{
		"message": "Lambda in GO for create units",
	})

	if err != nil {
		fmt.Println("Ingresa IF error")
		resp := ResponseUnit{
			StatusCode: 404,
			Body: fmt.Sprintf("%+v", types.ErrorResponse{
				ErrorCode:    001,
				ErrorMessage: "Check logs",
				Timestamp:    "20/07/2023",
			}),
		}
		return resp, err
	}
	json.HTMLEscape(&buf, body)
	fmt.Println("Pasa IF ERROR")
	resp := ResponseUnit{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Content-Type":           "application/json",
			"X-MyCompany-Func-Reply": "user-handler",
		},
	}
	fmt.Println("Antes del return")
	return resp, nil
}

func main() {
	lambda.Start(HandlerCreateUnit)
}
