package models

import (
	"time"

	database "github.com/marcio-olmedo-cavalini/financial-transactions-go-webapp/database"
	"gorm.io/gorm"
)

type TransactionReport struct {
	gorm.Model
	DataTransacao  time.Time `json:"dataTransacao"`
	DataImportacao time.Time `json:"dataImportacao"`
}

func GetAllTransactionReport() []TransactionReport {
	var transactionList []TransactionReport
	database.DB.Find(&transactionList).Order("data_importacao DESC")
	return transactionList
}

func CreateTransactionReport(transactionReport TransactionReport) {
	database.DB.Create(&transactionReport)
}
