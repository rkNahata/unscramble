package transactions

type Response struct {
	TransactionID       int     `json:"transactionId"`
	ProductName         string  `json:"productName"`
	TransactionAmount   float32 `json:"transactionAmount"`
	TransactionDateTime string  `json:"transactionDateTime"`
}

type ProductSummary map[string]float32

type TxnSummaryByProduct struct {
	Summary ProductSummary `json:"summary"`
}

type CitySummary map[string]float32
type TxnSummaryByCity struct{
	Summary CitySummary `json:"summary"`
}