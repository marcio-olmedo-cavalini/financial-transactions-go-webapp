package controllers

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/marcio-olmedo-cavalini/financial-transactions-go-webapp/helpers"
	models "github.com/marcio-olmedo-cavalini/financial-transactions-go-webapp/models"
)

func UploadFile(c *gin.Context) {
	file, err := c.FormFile("myFile")
	if err != nil {
		c.String(http.StatusBadRequest, "get form err: %s", err.Error())
		return
	}

	filename := filepath.Base(file.Filename)
	fmt.Println(filename)

	if err := c.SaveUploadedFile(file, "upload/"+filename); err != nil {
		c.String(http.StatusBadRequest, "upload file err: %s", err.Error())
		return
	}

	totalRows := readAndPrintUploadedFile(filename)
	validationMessage := ""
	if totalRows == 0 {
		validationMessage = "O arquivo informado está vazio"
	}

	validationMessage = readAndLoadUploadedFile(filename, helpers.GetLoggedUser(c))
	var transactionList = models.GetAllTransactionReportRawQuery()

	c.HTML(http.StatusOK, "upload.html", gin.H{
		"mensagem":         validationMessage,
		"arquivo":          file.Filename,
		"tamaho":           strconv.Itoa(int(file.Size)),
		"quantidadeLinhas": totalRows,
		"transactionList":  transactionList,
	})
}

func openUploadedFile(filename string) (http.File, [][]string) {
	d := http.Dir("./upload")
	f, err := d.Open(filename)
	if err != nil {
		panic(err)
	}

	filedata, err := csv.NewReader(f).ReadAll()
	if err != nil {
		log.Println(err)
		return nil, nil
	}

	return f, filedata
}

func readAndPrintUploadedFile(filename string) int {
	f, filedata := openUploadedFile(filename)
	totalRows := len(filedata)

	for e, value := range filedata {
		fmt.Println(e, value)
	}

	defer f.Close()
	io.Copy(os.Stdout, f)
	return totalRows
}

func readAndLoadUploadedFile(filename string, emailLoggedUser string) string {
	f, filedata := openUploadedFile(filename)
	var dataLote time.Time
	var dataTransacao time.Time

	//fmt.Println("logged user: " + emailLoggedUser)

	for e, value := range filedata {
		if e == 0 {
			dataLoteTmp, _ := time.Parse("2006-01-02T15:04:05", value[7])
			dataLote = time.Date(dataLoteTmp.Year(), dataLoteTmp.Month(), dataLoteTmp.Day(), 0, 0, 0, 0, dataLoteTmp.Location())

			if validateDateTransaction(dataLote) {
				fmt.Println("O Lote com essa data já está cadastrado!")
				return "O Lote com essa data já está cadastrado!"
			}
		}

		dataTransacao, _ = time.Parse("2006-01-02T15:04:05", value[7])
		dataTransacao = time.Date(dataTransacao.Year(), dataTransacao.Month(), dataTransacao.Day(), 0, 0, 0, 0, dataTransacao.Location())

		if dataTransacao == dataLote {
			var financialTransaction = new(models.FinancialTransaction)
			financialTransaction.BancoOrigem = value[0]
			financialTransaction.AgenciaOrigem = value[1]
			financialTransaction.ContaOrigem = value[2]
			financialTransaction.BancoDestino = value[3]
			financialTransaction.AgenciaDestino = value[4]
			financialTransaction.ContaDestino = value[5]
			financialTransaction.ValorTransacao, _ = strconv.ParseFloat(value[6], 64)

			dateString := value[7]
			dateConverted, _ := time.Parse("2006-01-02T15:04:05", dateString)
			financialTransaction.DataHoraTransacao = dateConverted
			//fmt.Println(e, financialTransaction.DataHoraTransacao)
			if err := models.ValidateFinancialTransaction(financialTransaction); err == nil {
				models.CreateFinancialTransaction(*financialTransaction)
			}
		}

		if e == len(filedata)-1 {
			user := models.FindUserByEmail(emailLoggedUser)
			transactionReport := models.TransactionReport{DataTransacao: dataLote, DataImportacao: time.Now(), UserImportacao: user}
			models.CreateTransactionReport(transactionReport)
		}
	}

	defer f.Close()
	return ""
}

func validateDateTransaction(dataLote time.Time) bool {
	return models.ExistsFinancialTransactionByDate(dataLote)
}

func ShowDetailImportPage(c *gin.Context) {
	id := c.Query("id")
	fmt.Println("id" + id)
	transactionReport := models.GetTransactionById(id)
	financialTransaction := models.GetAllFinancialTransactionRawQuery(id)
	c.HTML(http.StatusOK, "detailimport.html", gin.H{
		"DataImportacao": transactionReport.DataImportacao,
		"NomeUsuario":    transactionReport.NomeUsuario,
		"DataTransacao":  transactionReport.DataTransacao,
		"transactions":   financialTransaction,
	})
}
