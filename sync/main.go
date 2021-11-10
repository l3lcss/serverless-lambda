package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Handler func
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Setup
	// config := appconf.Read()
	// mySQLConn := database.ConnectMySQL(config.Mysql)
	// err := mySQLConn.Migrator().DropTable(models.Brand{}, models.BrandLanguage{}, models.Category{}, models.CategoryLanguage{}, models.Product{}, models.ProductLanguage{}, models.Sku{}, models.SkuLanguage{}, models.SellingPriceArea{}, models.SellingPrivatePrice{}, models.Image{}, models.Unit{}, models.UnitLanguage{})
	// if err != nil {
	// 	panic("DropTable failed.")
	// }

	envVariable1 := os.Getenv("variable1")
	envGOOS := os.Getenv("GOOS")
	fmt.Println("envenvVariable1: ", envVariable1)
	fmt.Println("envGOOS: ", envGOOS)
	fmt.Println("Received body: ", request.Body)

	return events.APIGatewayProxyResponse{Body: request.Body, StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}
