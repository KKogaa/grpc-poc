package clients

import (
	"context"

	"github.com/KKogaa/grpc-transaction/internal/adapters/outbound/grpc_pb"
	"github.com/KKogaa/grpc-transaction/internal/core/entities"
	"google.golang.org/grpc"
)

type NotificationClient struct {
	client grpc_pb.NotificationServiceClient
	conn   *grpc.ClientConn
}

func NewNotificationClient(address string) (*NotificationClient, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	client := grpc_pb.NewNotificationServiceClient(conn)

	return &NotificationClient{
		client: client,
		conn:   conn,
	}, nil
}

func (n *NotificationClient) SendNotification(transaction entities.Transaction) error {
	req := &grpc_pb.TransactionNotification{
		Id:          transaction.Id,
		Amount:      transaction.Amount,
		Description: transaction.Description,
		Status:      string(transaction.Status),
	}

	_, err := n.client.SendNotification(context.Background(), req)
	if err != nil {
		return err
	}

	return nil
}
