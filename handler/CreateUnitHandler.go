package main

import (
	"bytes"
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
func HandlerCreateUnit(event events.APIGatewayProxyRequest) (ResponseUnit, error) {
	fmt.Println("[INFO]HandlerCreateUnit init")
	fmt.Println("[INFO]Event: ", event.Body)
	var buf bytes.Buffer
	err := service.CreateUnit(event)

	if err != nil {
		fmt.Println("[Error]err: ", err.Error())
		resp := ResponseUnit{
			StatusCode: 404,
			Body: fmt.Sprintf("%+v", types.ErrorResponse{
				ErrorCode:    001,
				ErrorMessage: err.Error(),
				Timestamp:    "20/07/2023",
			}),
		}
		return resp, err
	}

	body, err := json.Marshal(map[string]interface{}{
		"message": "Unit created successful",
	})
	json.HTMLEscape(&buf, body)
	resp := ResponseUnit{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Content-Type":           "application/json",
			"X-MyCompany-Func-Reply": "user-handler",
		},
	}
	return resp, nil
}

func main() {
	lambda.Start(HandlerCreateUnit)
}
