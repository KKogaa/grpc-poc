package main

import (
	"log"
	"net"

	"github.com/KKogaa/grpc-transaction/internal/adapters/inbound/grpc_pb"
	"github.com/KKogaa/grpc-transaction/internal/adapters/inbound/handlers"
	"github.com/KKogaa/grpc-transaction/internal/adapters/outbound/repositories"
	"github.com/KKogaa/grpc-transaction/internal/core/services"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	db, err := sqlx.Connect("postgres", "postgresql://retool:npg_ti4HDmeYko6X@ep-blue-tree-a6w8ipns.us-west-2.retooldb.com/retool?sslmode=require")
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}

	grpcServer := grpc.NewServer()

	transactionRepository := repositories.NewTransactionRepository(db)
	transactionService := services.NewTransactionService(transactionRepository)
	transactionHandler := handlers.NewTransactionHandler(transactionService)

	grpc_pb.RegisterTransactionServiceServer(grpcServer,
		transactionHandler)
	reflection.Register(grpcServer)

	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("gRPC server is running on port :8080")

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
