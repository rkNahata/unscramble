package transactions

import (
	"fmt"
	"github.com/gocarina/gocsv"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"time"
	"unscramble/data"
)

var Service IService = service{}

type IService interface {
	GetByID(id int) (*Response, error)
	GetSummaryByProduct(days int) (*TxnSummaryByProduct, error)
	GetSummaryByCity(days int) (*TxnSummaryByCity, error)
}

type service struct{}

const DateTimeLayout = "2006-01-02 15:04:05"
const BaseTransaction = "transactionFiles"

func (service) GetByID(id int) (*Response, error) {
	var response Response
	td := GetTransactionsDetailsMap()
	pd := data.GetProductDetailsMap()
	if val, ok := td[id]; ok {
		response.TransactionID = id
		response.ProductName = pd[val.ProductID].ProductName
		response.TransactionAmount = val.TransactionAmount
		response.TransactionDateTime = val.TransactionDateTime.Format(DateTimeLayout)
	}
	return &response, nil

}

func (service) GetSummaryByProduct(days int) (*TxnSummaryByProduct, error) {
	currentTime := time.Now()
	var response TxnSummaryByProduct
	pds := make(ProductSummary)
	td := GetTransactionsDetailsMap()
	pd := data.GetProductDetailsMap()
	for _, t := range td {
		if t.TransactionDateTime.AddDate(0, 0, days).After(currentTime) {
			pds[pd[t.ProductID].ProductName] += t.TransactionAmount
		}
	}
	response.Summary = pds
	return &response, nil
}

func (service) GetSummaryByCity(days int) (*TxnSummaryByCity, error) {
	currentTime := time.Now()
	var response TxnSummaryByCity
	cds := make(CitySummary)
	td := GetTransactionsDetailsMap()
	pd := data.GetProductDetailsMap()
	for _, t := range td {
		if t.TransactionDateTime.AddDate(0, 0, days).After(currentTime) {
			cds[pd[t.ProductID].ProductManufacturingCity] += t.TransactionAmount
		}
	}
	response.Summary = cds
	return &response, nil
}

func getTransactionsFromCSV(filePaths []string) []*data.TransactionData {
	var transactionList []*data.TransactionData
	var tsf []*data.Transactions

	for _, file := range filePaths {
		var txnDataList []*data.TransactionData
		transactions, err := os.Open(file)
		if err != nil {
			fmt.Println(err)
			break
		}
		defer transactions.Close()
		err = gocsv.Unmarshal(transactions, &tsf)
		if err != nil {
			fmt.Println(err)
		}
		for _, t := range tsf {
			var tmp data.TransactionData
			if len(t.TransactionDateTime) > 0 {
				if validTime, err := time.Parse(DateTimeLayout, t.TransactionDateTime); err == nil {
					tmp.TransactionDateTime = &validTime
				}
			} else {
				fmt.Println("time format invalid")
				continue
			}
			tmp.TransactionID = t.TransactionID
			tmp.ProductID = t.ProductID
			tmp.TransactionAmount = t.TransactionAmount
			txnDataList = append(txnDataList, &tmp)
		}
		transactionList = append(transactionList, txnDataList...)
	}
	return transactionList
}

var transactionDetails map[int]*data.TransactionData

func CreateTransactionDetailsMapAtStartup() map[int]*data.TransactionData {
	var fileList []string
	wd, _ := os.Getwd()
	transactiondest := "transactionFiles"
	var err error
	filesInfoList, err := ioutil.ReadDir(transactiondest)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	for _, f := range filesInfoList {
		fileList = append(fileList, filepath.Join(wd, transactiondest, f.Name()))
	}
	td := GetTransactionsDetailsMap()
	txn := getTransactionsFromCSV(fileList)
	for _, t := range txn {
		td[t.TransactionID] = t
	}
	transactionDetails = td
	return transactionDetails
}

func GetTransactionsDetailsMap() map[int]*data.TransactionData {
	if transactionDetails != nil {
		return transactionDetails
	}
	return make(map[int]*data.TransactionData)
}

func CreateTransactionDetailsMap(filePath string) map[int]*data.TransactionData {
	filePaths := []string{path.Join(BaseTransaction, filePath)}
	td := GetTransactionsDetailsMap()
	txn := getTransactionsFromCSV(filePaths)
	for _, t := range txn {
		td[t.TransactionID] = t
	}
	transactionDetails = td
	return transactionDetails
}
