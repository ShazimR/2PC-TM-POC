package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"

	"google.golang.org/grpc"

	pb "2pc-tm-poc/transaction" // Update with your package name
)

// Placeholder functions to simulate rpc to other services
func SimulateService(name string, success bool) bool {
	var delay int = rand.Intn(5)
	fmt.Printf("Service %s is using the DB\n", name)
	time.Sleep(time.Duration(delay) * time.Second)

	if success {
		fmt.Printf("Service %s complete!\n", name)
	} else {
		fmt.Printf("Service %s aborted!\n", name)
	}

	return success
}

func isSuccessful(A bool, B bool, C bool) bool {
	return A && B && C
}

// TransactionManagerServer represents the transaction manager
type TransactionManagerServer struct {
	pb.UnimplementedTransactionManagerServer
}

// PerformOperation handles the PerformOperation RPC call
func (s *TransactionManagerServer) PerformOperation(ctx context.Context, req *pb.OperationRequest) (*pb.OperationResponse, error) {
	test := req.GetTest()
	A := test[0] == '1'
	B := test[1] == '1'
	C := test[2] == '1'

	outA := make(chan bool)
	outB := make(chan bool)
	outC := make(chan bool)

	go func() { outA <- SimulateService("A", A) }()
	go func() { outB <- SimulateService("B", B) }()
	go func() { outC <- SimulateService("C", C) }()

	success := isSuccessful(<-outA, <-outB, <-outC)

	message := "Operation Aborted!"
	if success {
		message = "Operation Commited!"
	}

	return &pb.OperationResponse{Success: success, Message: message}, nil
}

func main() {
	// Create a new gRPC server
	server := grpc.NewServer()

	// Initialize the transaction manager
	transactionManager := &TransactionManagerServer{}

	// Register the transaction manager with the gRPC server
	pb.RegisterTransactionManagerServer(server, transactionManager)

	// Start the gRPC server
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	fmt.Println("Starting Transaction Manager server...")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
