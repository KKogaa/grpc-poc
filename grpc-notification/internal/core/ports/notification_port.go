package ports

import "github.com/KKogaa/grpc-notification/internal/core/entities"

type NotificationPort interface {
	CreateNotification(entities.Notification) (entities.Notification, error)
}
