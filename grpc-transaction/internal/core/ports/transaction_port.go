package ports

import "github.com/KKogaa/grpc-transaction/internal/core/entities"

type TransactionPort interface {
	CreateTransaction(entities.Transaction) (entities.Transaction, error)
	UpdateTransaction(entities.Transaction) (entities.Transaction, error)
	FindById(id string) (entities.Transaction, error)
}
