package services

import (
	"github.com/KKogaa/grpc-notification/internal/core/entities"
	"github.com/KKogaa/grpc-notification/internal/core/ports"
	"github.com/google/uuid"
)

type NotificationService struct {
	notificationRepository ports.NotificationPort
	emailClient            ports.EmailPort
}

func NewNotificationService(notificationRepository ports.NotificationPort,
	emailClient ports.EmailPort) *NotificationService {
	return &NotificationService{
		notificationRepository: notificationRepository,
		emailClient:            emailClient,
	}
}

func (n *NotificationService) SendNotification(transactionId string,
	amount float32, description string, status string) (entities.Notification, error) {

	email := entities.Email{
		To:            "andreskoga@gmail.com",
		Subject:       "Transaction Notification",
		TransactionId: transactionId,
		Amount:        amount,
		Description:   description,
		Status:        status,
	}

	_, err := n.emailClient.SendEmail(email)
	if err != nil {
		return entities.Notification{}, err
	}

	notification := entities.Notification{
		Id:            uuid.New().String(),
		TransactionId: transactionId,
		Status:        "sent",
	}
	_, err = n.notificationRepository.CreateNotification(notification)

	if err != nil {
		return entities.Notification{}, err
	}

	return notification, nil
}
