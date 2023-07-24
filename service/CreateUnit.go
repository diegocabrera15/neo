package service

import (
	"context"
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

func CreateUnit(name string, ctx context.Context, event events.APIGatewayProxyRequest) {
	fmt.Println("Estamos en la creación Unidades: ", name)
	fmt.Println("Body: ", event.Body)

	var unit types.UnitDynamodb
	err := json.Unmarshal([]byte(event.Body), &unit)
	fmt.Println("Valores de ERR", err)
	if err != nil {
		fmt.Println("Error la estructura de User es incorrecta")
	}

	//////
	var TableName = os.Getenv("UNIT_DYNAMODB")

	unitMap, marshalErr := dynamodbattribute.MarshalMap(unit)

	fmt.Println("NOMBRE TABLA2", os.Getenv("UNIT_DYNAMODB"))
	if marshalErr != nil {
		fmt.Println("Failed to marshal to dynamo map")
		return marshalErr
	}

	dynamoSession := connection.createDynamoSession()
	if dynamoSession != nil {
		fmt.Println("dynamoSession Err: ", dynamoSession.Error())
		fmt.Println("dynamoSession http: ", http.StatusInternalServerError)
	}

	input := &dynamodb.PutItemInput{
		Item:      unitMap,
		TableName: aws.String(TableName),
	}

	_, writeErr := dynamoSession.PutItem(input)
	if writeErr != nil {
		fmt.Println("Failed to write to dynamo")
		return writeErr
	}
	return nil
}
