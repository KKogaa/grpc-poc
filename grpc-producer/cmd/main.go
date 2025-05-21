package main

import (
	"log"
	"net"

	"github.com/KKogaa/grpc-producer/internal/adapters/outbound/grpc_pb"
	"github.com/KKogaa/grpc-producer/internal/adapters/outbound/handlers"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	grpc_pb.RegisterTransactionServiceServer(grpcServer,
		handlers.NewTransactionHandler())
	reflection.Register(grpcServer)

	log.Println("gRPC server is running on port :8080")
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
