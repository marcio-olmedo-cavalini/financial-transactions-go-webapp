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
	database.DB.Exec("INSERT INTO users (created_at, updated_at, nome, email, password) select now() as created_at, now() as updated_at, 'Admin' as nome, 'admin@admin.com' as email, '$2a$14$JTMKeXZCQmXoRV/8nxJV1.QLQwDDREZP2cdJCVBax8TPcMlbPYIPy' as password where not exists (select id from users where nome = 'Admin')")
}
