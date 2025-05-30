package main

import (
	"fmt"
	"log"
	"net"

	"github.com/KKogaa/grpc-transaction/config"
	"github.com/KKogaa/grpc-transaction/internal/adapters/inbound/grpc_pb"
	"github.com/KKogaa/grpc-transaction/internal/adapters/inbound/handlers"
	"github.com/KKogaa/grpc-transaction/internal/adapters/outbound/clients"
	"github.com/KKogaa/grpc-transaction/internal/adapters/outbound/repositories"
	"github.com/KKogaa/grpc-transaction/internal/core/services"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	config := config.LoadConfig()

	db, err := sqlx.Connect("postgres", config.DatabaseUrl)
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}
	defer db.Close()

	grpcServer := grpc.NewServer()

	transactionRepository := repositories.NewTransactionRepository(db)
	noticationClient, err := clients.NewNotificationClient(config.NotificationServiceUrl)
	if err != nil {
		log.Fatalf("failed to create notification client: %v", err)
	}
	transactionService := services.NewTransactionService(transactionRepository, noticationClient)
	transactionHandler := handlers.NewTransactionHandler(transactionService)

	grpc_pb.RegisterTransactionServiceServer(grpcServer,
		transactionHandler)
	reflection.Register(grpcServer)

	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("gRPC server is listening on port %s", config.Port)

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
