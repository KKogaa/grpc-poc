package ports

import "github.com/KKogaa/grpc-transaction/internal/core/entities"

type NotificationPort interface {
	SendNotification(entities.Transaction) error
}
