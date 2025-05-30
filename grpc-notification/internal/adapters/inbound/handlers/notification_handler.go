package handlers

import (
	"context"
	"io"
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

func (n *NotificationHandler) StreamNotifications(stream grpc_pb.NotificationService_StreamNotificationsServer) error {
	log.Println("New streaming client connected")

	ctx := stream.Context()

	for {
		select {
		case <-ctx.Done():
			log.Println("Stream context cancelled")
			return ctx.Err()
		default:
		}

		req, err := stream.Recv()
		if err == io.EOF {
			log.Println("Client closed the stream")
			return nil
		}
		if err != nil {
			log.Printf("Error receiving from stream: %v", err)
			return status.Errorf(codes.Internal, "stream receive error: %v", err)
		}

		log.Printf("Processing streaming notification: ID=%s, Amount=%.2f", req.Id, req.Amount)

		notification, err := n.notificationService.SendNotification(
			req.Id, req.Amount, req.Description, req.Status)

		log.Println("Notification processed")

		var ack *grpc_pb.NotificationAcknowledgment
		if err != nil {
			log.Printf("Failed to process notification %s: %v", req.Id, err)
			ack = &grpc_pb.NotificationAcknowledgment{
				Id:             req.Id,
				NotificationId: "",
				Status:         "failed",
			}
		} else {
			ack = &grpc_pb.NotificationAcknowledgment{
				Id:             req.Id,
				NotificationId: notification.Id,
				Status:         notification.Status,
			}
		}

		if err := stream.Send(ack); err != nil {
			log.Printf("Error sending acknowledgment: %v", err)
			return status.Errorf(codes.Internal, "stream send error: %v", err)
		}

		log.Printf("Sent acknowledgment for notification %s", req.Id)
	}
}
