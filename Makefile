build: 
	protoc --go_out=. --go-grpc_out=. transaction.proto
	go build .