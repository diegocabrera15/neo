package service

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/go/neo/connection"
	"github.com/go/neo/types"
	"net/http"
	"os"
)

func CreateUnit(event events.APIGatewayProxyRequest) error {
	fmt.Println("Estamos en la creación Unidades")
	fmt.Println("Body: ", event.Body)

	unit := event.Body
	unit := types.UnitDynamodb{}
	err := json.Unmarshal([]byte(event.Body), &unit)

	fmt.Println("Valores de ERR", err)
	if err != nil {
		fmt.Println("Error la estructura de User es incorrecta")
	}

	var TableName = os.Getenv("UNIT_DYNAMODB")

	unitMap, marshalErr := dynamodbattribute.MarshalMap(unit)
	if marshalErr != nil {
		fmt.Println("Failed to marshal to dynamo map")
		return marshalErr
	}

	dynamoSession := connection.CreateDynamoSession()
	input := &dynamodb.PutItemInput{
		Item:      unitMap,
		TableName: aws.String(TableName),
	}

	_, writeErr := dynamoSession.PutItem(input)
	if writeErr != nil {
		fmt.Println("Failed to write to dynamo")
		fmt.Println("writeErr Err: ", writeErr.Error())
		fmt.Println("writeErr http: ", http.StatusInternalServerError)
		return writeErr
	}

	return nil
}
