package entities

import "time"

type Notification struct {
	Id            string    `json:"id" db:"id"`
	TransactionId string    `json:"transaction_id" db:"transaction_id"`
	Status        string    `json:"status" db:"status"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
}
