package connection

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	// "github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/go/neo/types"
	"os"
	// "strconv"
)

var TableName = os.Getenv("UNIT_DYNAMODB")

func SaveUnit(unit types.UnitDynamodb) error {
	unitMap, marshalErr := dynamodbattribute.MarshalMap(unit)

	fmt.Println("NOMBRE TABLA2", os.Getenv("UNIT_DYNAMODB"))
	if marshalErr != nil {
		fmt.Println("Failed to marshal to dynamo map")
		return marshalErr
	}

	dynamoSession := createDynamoSession()

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

func createDynamoSession() *dynamodb.DynamoDB {
	sess := session.Must(session.NewSessionWithOptions(
		session.Options{
			SharedConfigState: session.SharedConfigEnable,
		},
	))

	return dynamodb.New(sess)
}

func ReadUnit(ctx context.Context, event events.APIGatewayProxyRequest) ([]types.UnitDynamodb, error) {
	fmt.Println("Ingresa a ReadUnit")

	fmt.Println("Path vars: ", event.PathParameters["unitId"])
	var itemArray events.APIGatewayProxyResponse
	itemArray, _ = ListByUnitIds(event)
	fmt.Println("VALORES itemArray EN ReadUnit", itemArray)

	/*if err != nil {
		panic(fmt.Sprintf("Failed to find Item, %v", err))
	}

	// Make sure the Item isn't empty
	if len(items) == 0 {
		fmt.Println("Could not find units ", event.PathParameters["unitId"])
		//return events.APIGatewayProxyResponse{Body: event.Body, StatusCode: 500}, nil
	}

	// Log and return result
	stringItems := "["
	for i := 0; i < len(items); i++ {
		jsonItem, _ := json.Marshal(items[i])
		stringItems += string(jsonItem)
		if i != len(items)-1 {
			stringItems += ",\n"
		}
	}
	stringItems += "]\n"
	fmt.Println("Found items: ", stringItems)*/

	return nil, nil
}

func ListByUnitId(unitId string) ([]types.UnitDynamodb, error) {
	fmt.Println("Ingresa a ListByUnitIds")
	/*// Build the Dynamo client object
	sess := session.Must(session.NewSession())
	svc := dynamodb.New(sess)

	items := []types.UnitDynamodb{}

	// Create the Expression to fill the input struct with.
	unitIdInt, err := strconv.Atoi(unitId)
	fmt.Println("VALUE unitIdInt:", unitIdInt)
	filt := expression.Name("unitId").Equal(expression.Value(unitIdInt))
	fmt.Println("VALUE filt:", filt)
	// Get back the title, year, and rating
	proj := expression.NamesList(expression.Name("unitId"))
	fmt.Println("VALUE proj:", proj)
	expr, err := expression.NewBuilder().WithFilter(filt).WithProjection(proj).Build()
	fmt.Println("VALUE expr:", expr)
	fmt.Println("VALUE TableName:", TableName)
	if err != nil {
		fmt.Println("Got error building expression:")
		fmt.Println(err.Error())
		return items, err
	}

	// Build the query input parameters
	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String(TableName),
	}

	// Make the DynamoDB Query API call
	result, err := svc.Scan(params)
	fmt.Println("Result", result)

	if err != nil {
		fmt.Println("Query API call failed:")
		fmt.Println((err.Error()))
		return items, err
	}

	numItems := 0
	for _, i := range result.Items {
		item := types.UnitDynamodb{}

		err = dynamodbattribute.UnmarshalMap(i, &item)

		if err != nil {
			fmt.Println("Got error unmarshalling:")
			fmt.Println(err.Error())
			return items, err
		}

		fmt.Println("Value record: ", item)
		items = append(items, item)
		numItems++
	}

	fmt.Println("Found", numItems, "UNITS", unitId)
	if err != nil {
		fmt.Println(err.Error())
		return items, err
	}*/

	return nil, nil
}

func ListByUnitIds(event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Ingresa a ListByUnitIds")
	// Build the Dynamo client object
	sess := session.Must(session.NewSession())
	svc := dynamodb.New(sess)

	// Build the query input parameters
	params := &dynamodb.ScanInput{
		TableName: aws.String(TableName),
	}

	// Scan table
	result, err := svc.Scan(params)

	// Checking for errors, return error
	if err != nil {
		fmt.Println("Query API call failed: ", err.Error())
		return events.APIGatewayProxyResponse{StatusCode: 500}, nil
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
			return events.APIGatewayProxyResponse{StatusCode: 500}, nil
		}

		itemArray = append(itemArray, item)
	}

	fmt.Println("itemArray: ", itemArray)
	itemArrayString, err := json.Marshal(itemArray)
	if err != nil {
		fmt.Println("Got error marshalling result: ", err.Error())
		return events.APIGatewayProxyResponse{StatusCode: 500}, nil
	}

	return events.APIGatewayProxyResponse{Body: string(itemArrayString), StatusCode: 200}, nil
}
