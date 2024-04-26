package web

type TransactionRequest struct {
	Description string  `json:"description"`
	Amount      float64 `json:"amount"`
}
