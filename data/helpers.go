package data

import (
	"fmt"
	"github.com/gocarina/gocsv"
	"os"
)

func getProductFromCSV(filePath string) []*Product {
	var products []*Product
	prd, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer prd.Close()
	err = gocsv.Unmarshal(prd, &products)
	if err != nil {
		fmt.Println(err)
	}
	return products
}

var productDetails map[int]*Product

func CreateProductDetailsMap() map[int]*Product {
	filePath := "product.csv"
	pd := make(map[int]*Product)
	products := getProductFromCSV(filePath)
	for _, p := range products {
		pd[p.ProductID] = p
	}
	productDetails = pd
	return productDetails
}

func GetProductDetailsMap() map[int]*Product {
	if productDetails != nil {
		return productDetails
	}
	return make(map[int]*Product)
}

var cityWiseProducts map[string][]int

func createCityWiseProductMap(filePath string) map[string][]int {
	c := make(map[string][]int)
	products := getProductFromCSV(filePath)
	for _, p := range products {
		c[p.ProductManufacturingCity] = append(c[p.ProductManufacturingCity], p.ProductID)
	}
	cityWiseProducts = c
	return cityWiseProducts
}

func GetCityWiseProductMap(filePath string) map[string][]int {
	if cityWiseProducts != nil {
		return cityWiseProducts
	}
	return createCityWiseProductMap(filePath)
}
