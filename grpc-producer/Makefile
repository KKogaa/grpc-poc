run:
	go run cmd/main.go
gen:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative internal/adapters/outbound/grpc_pb/*.proto
