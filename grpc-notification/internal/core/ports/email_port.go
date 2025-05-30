package ports

import "github.com/KKogaa/grpc-notification/internal/core/entities"

type EmailPort interface {
	SendEmail(entities.Email) (entities.Email, error)
}
