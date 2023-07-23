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

// ResponseReadUnit is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type ResponseReadUnit events.APIGatewayProxyResponse

// HandlerReadUnit is our lambda handler invoked by the `lambda.Start` function call
func HandlerReadUnit(ctx context.Context, event events.APIGatewayProxyRequest) (ResponseReadUnit, error) {
	fmt.Println("Ingresa a HandlerReadUnit")
	var buf bytes.Buffer

	service.ReadUnit(ctx, event)

	body, err := json.Marshal(map[string]interface{}{
		"message": "Lambda in GO for read units",
	})

	if err != nil {
		fmt.Println("Ingresa IF error")
		resp := ResponseReadUnit{
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
	resp := ResponseReadUnit{
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
	lambda.Start(HandlerReadUnit)
}
