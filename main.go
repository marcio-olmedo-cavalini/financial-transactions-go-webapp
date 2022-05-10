package main

import (
	"fmt"

	"github.com/marcio-olmedo-cavalini/financial-transactions-go-webapp/database"
	"github.com/marcio-olmedo-cavalini/financial-transactions-go-webapp/models"
	"github.com/marcio-olmedo-cavalini/financial-transactions-go-webapp/routes"
)

func main() {
	fmt.Println("Hello World!")
	database.OpenConnection()
	Migrate()
	routes.HandleRequests()

}

func Migrate() {
	database.DB.AutoMigrate(&models.FinancialTransaction{})
	database.DB.AutoMigrate(&models.TransactionReport{})
	database.DB.AutoMigrate(&models.User{})
}
