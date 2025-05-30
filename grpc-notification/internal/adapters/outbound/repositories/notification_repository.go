package repositories

import (
	"github.com/KKogaa/grpc-notification/internal/core/entities"
	"github.com/jmoiron/sqlx"
)

type NotificationRepository struct {
	db *sqlx.DB
}

func NewNotificationRepository(db *sqlx.DB) *NotificationRepository {
	return &NotificationRepository{
		db: db,
	}
}

func (r *NotificationRepository) CreateNotification(notification entities.Notification) (entities.Notification, error) {
	query := `INSERT INTO notifications (id, transaction_id, status) 
            VALUES (:id, :transaction_id, :status)`

	_, err := r.db.NamedExec(query, notification)
	if err != nil {
		return entities.Notification{}, err
	}

	return notification, nil
}
