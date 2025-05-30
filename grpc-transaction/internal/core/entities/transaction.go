package entities

import "time"

type Transaction struct {
	Id          string            `json:"id" db:"id"`
	Amount      float32           `json:"amount" db:"amount"`
	UserId      string            `json:"user_id" db:"user_id"`
	Description string            `json:"description" db:"description"`
	Status      TransactionStatus `json:"status" db:"status"`
	CreatedAt   time.Time         `json:"created_at" db:"created_at"`
}
