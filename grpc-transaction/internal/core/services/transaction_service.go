package services

import (
	"github.com/KKogaa/grpc-transaction/internal/core/entities"
	"github.com/KKogaa/grpc-transaction/internal/core/ports"
	"github.com/google/uuid"
)

type TransactionService struct {
	transactionRepository ports.TransactionPort
	notificationClient    ports.NotificationPort
}

func NewTransactionService(transactionRepository ports.TransactionPort,
	notificationClient ports.NotificationPort) *TransactionService {
	return &TransactionService{
		transactionRepository: transactionRepository,
		notificationClient:    notificationClient,
	}
}

func (t *TransactionService) CreateTransaction(amount float32,
	description string, userId string) (entities.Transaction, error) {

	transaction := entities.Transaction{
		Id:          uuid.New().String(),
		Amount:      amount,
		UserId:      uuid.New().String(),
		Description: description,
		Status:      entities.Pending,
	}

	_, err := t.transactionRepository.CreateTransaction(transaction)
	if err != nil {
		return entities.Transaction{}, err
	}

	return transaction, nil

}

func (t *TransactionService) UpdateTransactionStatus(id string,
	status string) (entities.Transaction, error) {

	transaction, err := t.transactionRepository.FindById(id)
	if err != nil {
		return entities.Transaction{}, err
	}

	transaction.Status = entities.TransactionStatus(status)

	_, err = t.transactionRepository.UpdateTransaction(transaction)
	if err != nil {
		return entities.Transaction{}, err
	}

	if err := t.notificationClient.SendNotification(transaction); err != nil {
		return entities.Transaction{}, err
	}

	return transaction, nil

}
