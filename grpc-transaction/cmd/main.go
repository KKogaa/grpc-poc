package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	db := setupDatabase(config)
	defer db.Close()

	transactionHandler := setupTransactionHandler(db, config)

	grpcServer := setupGRPCServer(transactionHandler)

	startServer(grpcServer, config.Port)

}

func setupDatabase(config *config.Config) *sqlx.DB {
	db, err := sqlx.Connect("postgres", config.DatabaseUrl)
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}
	return db
}

func setupTransactionHandler(db *sqlx.DB, config *config.Config) *handlers.TransactionHandler {
	notificationClient, err := clients.NewNotificationClient(config.NotificationServiceUrl)
	if err != nil {
		log.Fatalf("failed to create notification client: %v", err)
	}

	transactionRepo := repositories.NewTransactionRepository(db)
	transactionService := services.NewTransactionService(transactionRepo, notificationClient)
	return handlers.NewTransactionHandler(transactionService)
}

func setupGRPCServer(handler *handlers.TransactionHandler) *grpc.Server {
	server := grpc.NewServer()
	grpc_pb.RegisterTransactionServiceServer(server, handler)
	reflection.Register(server)
	return server
}

func startServer(server *grpc.Server, port string) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("failed to listen on port %s: %v", port, err)
		return
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Printf("gRPC server is listening on port %s", port)
		if err := server.Serve(listener); err != nil {
			log.Printf("failed to serve: %v", err)
		}
	}()

	<-stop
	log.Println("Shutting down gRPC server...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	done := make(chan struct{})
	go func() {
		server.GracefulStop()
		close(done)
	}()

	select {
	case <-done:
		log.Println("gRPC server stopped gracefully")
	case <-ctx.Done():
		log.Println("Graceful shutdown timeout, forcing stop...")
		server.Stop()
	}
}
