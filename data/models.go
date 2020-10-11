package data

import "time"

type Product struct {
	ProductID                int    `csv:"productId"`
	ProductName              string `csv:"productName"`
	ProductManufacturingCity string `csv:"productManufacturingCity"`
}

type Transactions struct {
	TransactionID       int     `csv:"transactionId"`
	ProductID           int     `csv:"productId"`
	TransactionAmount   float32 `csv:"transactionAmount"`
	TransactionDateTime string  `csv:"transactionDatetime"`
}

type ProductData struct {
	ProductID                int
	ProductName              string
	ProductManufacturingCity string
}

type TransactionData struct {
	TransactionID       int
	ProductID           int
	TransactionAmount   float32
	TransactionDateTime *time.Time
}
