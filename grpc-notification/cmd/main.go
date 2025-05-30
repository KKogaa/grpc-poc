package main

import (
	"log"
	"net"

	"github.com/KKogaa/grpc-notification/internal/adapters/inbound/grpc_pb"
	"github.com/KKogaa/grpc-notification/internal/adapters/inbound/handlers"
	"github.com/KKogaa/grpc-notification/internal/adapters/outbound/clients"
	"github.com/KKogaa/grpc-notification/internal/adapters/outbound/repositories"
	"github.com/KKogaa/grpc-notification/internal/core/services"
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
	defer db.Close()

	grpcServer := grpc.NewServer()

	emailClient := clients.NewEmailClient("https://api.retool.com/v1/workflows/87f81514-e028-4bbd-9966-00cf5cb4dd57/startTrigger",
		"retool_wk_17b6a986704248b4a171513a0d1085b1")
	notificationRepository := repositories.NewNotificationRepository(db)
	notificationService := services.NewNotificationService(notificationRepository, emailClient)
	notificationHandler := handlers.NewNotificationHandler(notificationService)

	grpc_pb.RegisterNotificationServiceServer(grpcServer,
		notificationHandler)
	reflection.Register(grpcServer)

	listen, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("gRPC server is running on port :8081")

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
