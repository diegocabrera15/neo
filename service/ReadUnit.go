package service

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/go/neo/constants"
	"github.com/go/neo/types"
)

func ReadUnit() events.APIGatewayProxyResponse {
	fmt.Println("[INFO]ReadUnit")
	var itemArray events.APIGatewayProxyResponse
	itemArray = ListByUnitIds()
	fmt.Println("[INFO]Valores itemArray", itemArray)
	return itemArray
}

func ListByUnitIds() events.APIGatewayProxyResponse {
	fmt.Println("Ingresa a ListByUnitIds")
	// Build the Dynamo client object
	sess := session.Must(session.NewSession())
	svc := dynamodb.New(sess)

	// Build the query input parameters
	params := &dynamodb.ScanInput{
		TableName: aws.String(constants.TableUnitDynamodb),
	}

	// Scan table
	result, err := svc.Scan(params)

	// Checking for errors, return error
	if err != nil {
		fmt.Println("Query API call failed: ", err.Error())
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body: fmt.Sprintf("%+v", types.ErrorResponse{
				ErrorCode:    001,
				ErrorMessage: err.Error(),
				Timestamp:    "20/07/2023",
			}),
		}
	}

	var itemArray []types.UnitDynamodb

	for _, i := range result.Items {
		item := types.UnitDynamodb{}

		// result is of type *dynamodb.GetItemOutput
		// result.Item is of type map[string]*dynamodb.AttributeValue
		// UnmarshallMap result.item to item
		err = dynamodbattribute.UnmarshalMap(i, &item)

		if err != nil {
			fmt.Println("Got error unmarshalling: ", err.Error())
			return events.APIGatewayProxyResponse{
				StatusCode: 500,
				Body: fmt.Sprintf("%+v", types.ErrorResponse{
					ErrorCode:    001,
					ErrorMessage: err.Error(),
					Timestamp:    "20/07/2023",
				}),
			}
		}
		itemArray = append(itemArray, item)
	}

	fmt.Println("itemArray: ", itemArray)
	itemArrayString, err := json.Marshal(itemArray)
	if err != nil {
		fmt.Println("Got error marshalling result: ", err.Error())
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body: fmt.Sprintf("%+v", types.ErrorResponse{
				ErrorCode:    001,
				ErrorMessage: err.Error(),
				Timestamp:    "20/07/2023",
			}),
		}
	}

	return events.APIGatewayProxyResponse{Body: string(itemArrayString), StatusCode: 200}
}
