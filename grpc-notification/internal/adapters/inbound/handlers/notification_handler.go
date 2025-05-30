package handlers

import (
	"context"
	"log"

	"github.com/KKogaa/grpc-notification/internal/adapters/inbound/grpc_pb"
	"github.com/KKogaa/grpc-notification/internal/core/services"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type NotificationHandler struct {
	grpc_pb.UnimplementedNotificationServiceServer
	notificationService *services.NotificationService
}

func NewNotificationHandler(notificationService *services.NotificationService) *NotificationHandler {
	return &NotificationHandler{
		notificationService: notificationService,
	}
}

func (n *NotificationHandler) SendNotification(ctx context.Context,
	req *grpc_pb.TransactionNotification) (*grpc_pb.NotificationAcknowledgment, error) {
	log.Printf("received request: %v", req)

	notification, err := n.notificationService.SendNotification(req.Id,
		req.Amount, req.Description, req.Status)

	if err != nil {
		log.Println("error: ", err)
		return nil, status.Errorf(codes.Internal, "failed to send notification")
	}

	return &grpc_pb.NotificationAcknowledgment{
		Id:     notification.Id,
		Status: notification.Status,
	}, nil
}
