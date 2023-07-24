package service

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/go/neo/connection"
	"github.com/go/neo/constants"
	"github.com/go/neo/types"
	"net/http"
)

func CreateUnit(event events.APIGatewayProxyRequest) error {
	fmt.Println("[INFO]CreateUnit")
	unit := types.UnitDynamodb{}
	err := json.Unmarshal([]byte(event.Body), &unit)
	if err != nil {
		fmt.Println("[ERROR]Error la estructura de User es incorrecta", err)
	}

	//if unit.UnitId == "" {
	//	return events.APIGatewayProxyResponse{StatusCode: 400}, nil
	//}

	item := types.UnitDynamodb{
		UnitId:      unit.UnitId,
		Name:        unit.Name,
		Unit:        unit.Unit,
		Description: unit.Description,
		State:       unit.State,
	}

	unitMap, marshalErr := dynamodbattribute.MarshalMap(item)
	if marshalErr != nil {
		fmt.Println("[ERROR]Failed to marshal to dynamo map")
		return marshalErr
	}

	dynamoSession := connection.CreateDynamoSession()
	input := &dynamodb.PutItemInput{
		Item:      unitMap,
		TableName: aws.String(constants.TableUnitDynamodb),
	}

	_, writeErr := dynamoSession.PutItem(input)
	if writeErr != nil {
		fmt.Println("[ERROR]Failed to write to dynamo")
		fmt.Println("[ERROR]WriteErr Err: ", writeErr.Error())
		fmt.Println("[ERROR]writeErr http: ", http.StatusInternalServerError)
		return writeErr
	}
	fmt.Println("[INFO]Unit created successful")

	return nil
}
