package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/go/neo/connection"
	"github.com/go/neo/types"
	"net/http"
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

	conn := connection.SaveUnit(unit)
	if conn != nil {
		fmt.Println("dynamoErr: ", conn.Error())
		fmt.Println("http: ", http.StatusInternalServerError)
	}
}
