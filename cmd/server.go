package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/radovskyb/watcher"
	"os"
	"os/signal"
	"strconv"
	"time"
	"unscramble/api"
	"unscramble/api/transactions"
	"unscramble/data"
	intTxn "unscramble/internal/transactions"
)

const Port = 8080

func main() {
	//initialize product details
	data.CreateProductDetailsMap()
	//initialize txn details
	intTxn.CreateTransactionDetailsMapAtStartup()
	fmt.Println("Starting server")
	//initialize the server
	startHttpServer()
	//start the transactionFiles dir watcher
	setUpWatcher()
	shutDownChan := make(chan os.Signal)
	signal.Notify(shutDownChan, os.Interrupt)
	<-shutDownChan
}

func startHttpServer() {
	router := gin.New()

	router.Use(
		gin.Recovery(),
	)
	assignmentGroup := router.Group("/assignment")
	assignmentGroup.GET("/transaction/:transaction_id", api.CommonHandler(transactions.Handler))
	assignmentGroup.GET("/transactionSummaryByProducts/:last_n_days", api.CommonHandler(transactions.ProductWiseSummaryHandler))
	assignmentGroup.GET("/transactionSummaryByManufacturingCity/:last_n_days", api.CommonHandler(transactions.CityWiseSummaryHandler))

	go func() {
		fmt.Print("Starting Web server port ", Port)
		if err := router.Run(":" + strconv.Itoa(Port)); err != nil {
			os.Exit(1)
		}
	}()

}

func setUpWatcher(){
	w := watcher.New()
	w.FilterOps(watcher.Create)
	if err := w.Add("./transactionFiles"); err != nil {
		fmt.Println(err)
	}
	go func() {
		for {
			select {
			case event := <-w.Event:
				intTxn.CreateTransactionDetailsMap(event.FileInfo.Name())
			case err := <-w.Error:
				fmt.Println(err)
			case <-w.Closed:
				return
			}
		}
	}()

	if err := w.Start(time.Millisecond * 100); err != nil {
		fmt.Println(err)
	}
}

