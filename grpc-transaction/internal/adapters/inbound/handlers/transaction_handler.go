package handlers

import (
	"context"
	"errors"
	"log"

	"github.com/KKogaa/grpc-transaction/internal/adapters/inbound/grpc_pb"
	customerrors "github.com/KKogaa/grpc-transaction/internal/core/custom_errors"
	"github.com/KKogaa/grpc-transaction/internal/core/services"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TransactionHandler struct {
	grpc_pb.UnimplementedTransactionServiceServer
	transactionService *services.TransactionService
}

func NewTransactionHandler(transactionService *services.TransactionService) *TransactionHandler {
	return &TransactionHandler{
		transactionService: transactionService,
	}
}

func (t *TransactionHandler) CreateTransaction(ctx context.Context,
	req *grpc_pb.CreateTransactionRequest) (*grpc_pb.CreateTransactionResponse, error) {
	log.Printf("received request: %v", req)

	transaction, err := t.transactionService.CreateTransaction(req.Amount, req.Description, "test")
	if err != nil {
		log.Println("error: ", err)
	}

	return &grpc_pb.CreateTransactionResponse{
		Id:          transaction.Id,
		Amount:      transaction.Amount,
		Description: transaction.Description,
		Status:      string(transaction.Status),
	}, nil
}

func (t *TransactionHandler) UpdateTransactionStatus(ctx context.Context,
	req *grpc_pb.UpdateTransactionStatusRequest) (*grpc_pb.UpdateTransactionStatusResponse, error) {
	log.Printf("received request: %v", req)

	transaction, err := t.transactionService.UpdateTransactionStatus(req.Id, req.Status)
	if err != nil {
		switch {
		case errors.Is(err, customerrors.ErrTransactionNotFound):
			log.Printf("transaction with id %s not found", req.Id)
			return nil, status.Errorf(codes.NotFound, "transaction not found")
		case errors.Is(err, customerrors.ErrTransactionRepository):
			log.Println("transaction repository error: ", err)
			return nil, status.Errorf(codes.Internal, "unexpected error")
		default:
			log.Println("unexpected error: ", err)
			return nil, status.Errorf(codes.Internal, "unexpected error")
		}
	}

	return &grpc_pb.UpdateTransactionStatusResponse{
		Id:          transaction.Id,
		Amount:      transaction.Amount,
		Description: transaction.Description,
		Status:      string(transaction.Status),
	}, nil
}
