gen-proto:
	protoc --go-grpc_out=. --go_out=. pb/*.proto 

go-run:
	go run cmd/main.go
