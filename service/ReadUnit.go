package service

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/go/neo/connection"
	"github.com/go/neo/types"
	"net/http"
)

func ReadUnit(ctx context.Context, event events.APIGatewayProxyRequest) ([]types.UnitDynamodb, error) {
	fmt.Println("Estamos en la Lectura de unidades")

	units, err := connection.ReadUnit(ctx, event)
	fmt.Println("UNIDADES LEIDAS EN LA TABLA", units)

	if units != nil {
		fmt.Println("dynamoErr1: ", units)
		fmt.Println("dynamoErr2: ", err)
		fmt.Println("http: ", http.StatusInternalServerError)
	}

	return units, nil
}
