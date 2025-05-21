package handlers

import (
	"context"
	"log"

	"github.com/KKogaa/grpc-producer/internal/adapters/outbound/grpc_pb"
)

type TransactionHandler struct {
	grpc_pb.UnimplementedTransactionServiceServer
}

func NewTransactionHandler() *TransactionHandler {
	return &TransactionHandler{}
}

func (t *TransactionHandler) CreateTransaction(ctx context.Context,
	req *grpc_pb.CreateTransactionRequest) (*grpc_pb.CreateTransactionResponse, error) {
	log.Printf("received request: %v", req)
	return &grpc_pb.CreateTransactionResponse{
		Id:     req.Id,
		UserId: req.UserId,
		Amount: req.Amount,
	}, nil
}
