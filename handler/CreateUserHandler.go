package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/go/neo/service"
	"github.com/go/neo/types"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
//type Response events.APIGatewayProxyResponse

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context) (types.ApiResponse, error) {
	var buf bytes.Buffer
	fmt.Println("Antes del Archivo")
	service.CreateUser("Diego")
	fmt.Println("Después del Archivo")
	body, err := json.Marshal(map[string]interface{}{
		"message": "My first lambda in GO",
	})

	if err != nil {
		fmt.Println("Ingresa IF error")
		resp := types.ApiResponse{
			StatusCode: 404,
			//Body:       err.Error(),
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
	resp := types.ApiResponse{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Content-Type":           "application/json",
			"X-MyCompany-Func-Reply": "user-handler",
		},
		//Errors: types.ErrorResponse{},
	}
	fmt.Println("Antes del return")
	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
