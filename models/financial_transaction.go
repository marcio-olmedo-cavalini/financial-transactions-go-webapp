package models

import (
	"time"

	database "github.com/marcio-olmedo-cavalini/financial-transactions-go-webapp/database"
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type FinancialTransaction struct {
	gorm.Model
	BancoOrigem       string    `json:"bancoOrigem" validate:"nonzero"`
	AgenciaOrigem     string    `json:"agenciaOrigem" validate:"nonzero"`
	ContaOrigem       string    `json:"contaOrigem" validate:"nonzero"`
	BancoDestino      string    `json:"bancoDestino" validate:"nonzero"`
	AgenciaDestino    string    `json:"agenciaDestino" validate:"nonzero"`
	ContaDestino      string    `json:"contaDestino" validate:"nonzero"`
	ValorTransacao    float64   `json:"valorTransacao" validate:"nonzero"`
	DataHoraTransacao time.Time `json:"dataHoraTransacao" validate:"nonzero"`
}

func ValidateFinancialTransaction(financialTransaction *FinancialTransaction) error {
	if err := validator.Validate(financialTransaction); err != nil {
		return err
	}

	return nil
}

func CreateFinancialTransaction(financialTransaction FinancialTransaction) {
	database.DB.Create(&financialTransaction)
}

func ExistsFinancialTransactionByDate(dataFiltro time.Time) bool {
	var ft FinancialTransaction
	result := database.DB.Where("data_hora_transacao >= ?", dataFiltro.Format("2006-01-02")).First(&ft)
	return result.RowsAffected != 0
}
