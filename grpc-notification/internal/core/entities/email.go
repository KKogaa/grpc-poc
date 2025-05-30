package entities

type Email struct {
	To            string  `json:"to"`
	Subject       string  `json:"subject"`
	TransactionId string  `json:"transaction_id"`
	Amount        float32 `json:"amount"`
	Description   string  `json:"description"`
	Status        string  `json:"status"`
}
