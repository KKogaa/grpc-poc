package handlers

import (
	"context"
	"log"

	"github.com/KKogaa/grpc-producer/internal/adapters/outbound/grpc_pb"
)

type InventoryHandler struct {
	grpc_pb.UnimplementedInventoryServiceServer
}

func NewInventoryHandler() *InventoryHandler {
	return &InventoryHandler{}
}

func (t *InventoryHandler) GetInventory(ctx context.Context,
	req *grpc_pb.CreateTransactionRequest) (*grpc_pb.CreateTransactionResponse, error) {
	log.Printf("received request: %v", req)
	return &grpc_pb.CreateTransactionResponse{
		Id:     req.Id,
		UserId: req.UserId,
		Amount: req.Amount,
	}, nil
}
