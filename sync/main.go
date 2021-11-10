package main

import (
	"fmt"
	"myservice/sync/appconf"
	"myservice/sync/database"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Handler func
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Setup
	config := appconf.Read()
	fmt.Printf("\n config = %+v \n\n", config)

	mySQLConn := database.ConnectMySQL(config.Mysql)
	sqlDB, err := mySQLConn.DB()
	if err != nil {
		panic(err)
	}
	defer sqlDB.Close()

	fmt.Println("mySQLConn: ", mySQLConn)
	fmt.Println("Received body: ", request.Body)

	return events.APIGatewayProxyResponse{Body: request.Body, StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}
