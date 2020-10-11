package transactions

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"unscramble/api"
	"unscramble/internal/transactions"
)

var Handler api.Handler = handler{transactions.Service}
var ProductWiseSummaryHandler api.Handler = productSummaryHandler{transactions.Service}
var CityWiseSummaryHandler api.Handler = citySummaryHandler{transactions.Service}

type handler struct {
	service transactions.IService
}

func (handler)CreateRequest(c *gin.Context)(interface{},error){
	return Request{},nil
}

func (handler)Handle(c *gin.Context,req interface{}) (interface{}, error) {
	txnId, err := strconv.Atoi(c.Param("transaction_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "InvalidRequest",
		})
		return nil, err
	}

	resp, err := transactions.Service.GetByID(txnId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "InvalidRequest",
		})
		return nil, err
	}
	return resp, nil

}

func (handler)CreateResponse(resp interface{})(interface{},error){
	return Response{},nil
}


type productSummaryHandler struct{
	service transactions.IService
}

func (productSummaryHandler)CreateRequest(c *gin.Context)(interface{},error){
	return Request{},nil
}

func (productSummaryHandler)Handle(c *gin.Context,req interface{}) (interface{}, error) {
	days, err := strconv.Atoi(c.Param("last_n_days"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "InvalidRequest",
		})
		return nil, err
	}

	resp, err := transactions.Service.GetSummaryByProduct(days)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "InvalidRequest",
		})
		return nil, err
	}
	return resp, nil

}

func (productSummaryHandler)CreateResponse(resp interface{})(interface{},error){
	return Response{},nil
}


type citySummaryHandler struct{
	service transactions.IService
}

func (citySummaryHandler)CreateRequest(c *gin.Context)(interface{},error){
	return Request{},nil
}

func (citySummaryHandler)Handle(c *gin.Context,req interface{}) (interface{}, error) {
	days, err := strconv.Atoi(c.Param("last_n_days"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "InvalidRequest",
		})
		return nil, err
	}

	resp, err := transactions.Service.GetSummaryByCity(days)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "InvalidRequest",
		})
		return nil, err
	}
	return resp, nil

}

func (citySummaryHandler)CreateResponse(resp interface{})(interface{},error){
	return Response{},nil
}