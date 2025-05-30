package entities

type TransactionStatus string

const (
	Pending    TransactionStatus = "pending"
	Processing TransactionStatus = "processing"
	Completed  TransactionStatus = "completed"
	Failed     TransactionStatus = "failed"
	Cancelled  TransactionStatus = "cancelled"
	Refunded   TransactionStatus = "refunded"
)
