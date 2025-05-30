package repositories

import (
	"fmt"
	"log"

	customerrors "github.com/KKogaa/grpc-transaction/internal/core/custom_errors"
	"github.com/KKogaa/grpc-transaction/internal/core/entities"
	"github.com/jmoiron/sqlx"
)

type TransactionRepository struct {
	db *sqlx.DB
}

func NewTransactionRepository(db *sqlx.DB) *TransactionRepository {
	return &TransactionRepository{
		db: db,
	}
}

func (t *TransactionRepository) CreateTransaction(transaction entities.Transaction) (entities.Transaction, error) {
	query := `
    INSERT INTO transactions (id, amount, user_id, description, status)
    VALUES (:id, :amount, :user_id, :description, :status)
  `

	_, err := t.db.NamedExec(query, &transaction)
	if err != nil {
		return entities.Transaction{}, customerrors.ErrTransactionRepository
	}

	return transaction, nil
}

func (t *TransactionRepository) UpdateTransaction(transaction entities.Transaction) (entities.Transaction, error) {
	query := `
        UPDATE transactions 
        SET amount = :amount, 
            user_id = :user_id, 
            description = :description, 
            status = :status
        WHERE id = :id`

	result, err := t.db.NamedExec(query, &transaction)
	if err != nil {
		return entities.Transaction{}, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return entities.Transaction{}, customerrors.ErrTransactionRepository
	}

	if rowsAffected == 0 {
		return entities.Transaction{}, fmt.Errorf("transaction with id %s not found", transaction.Id)
	}

	return transaction, nil
}

func (t *TransactionRepository) FindById(id string) (entities.Transaction, error) {
	var transaction entities.Transaction

	query := "SELECT * FROM transactions WHERE id = $1"
	err := t.db.Get(&transaction, query, id)
	if err != nil {
		log.Printf("error finding transaction by id %s: %v", id, err)
		return entities.Transaction{}, customerrors.ErrTransactionNotFound
	}

	return transaction, nil
}
