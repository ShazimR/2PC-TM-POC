package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"

	"google.golang.org/grpc"

	pb "2pc-tm-poc/transaction"
)

// Placeholder functions to simulate rpc to other services
func PrepareSimulateService(name string, success bool) bool {
	var delay int = rand.Intn(5)
	fmt.Printf("Service %s is using the DB\n", name)
	time.Sleep(time.Duration(delay) * time.Second)

	if success {
		fmt.Printf("Service %s completed!\n", name)
	} else {
		fmt.Printf("Service %s failed!\n", name)
	}

	return success
}

func CommitSimulateService(name string) {
	fmt.Printf("Service %s commited\n", name)
}

func AbortSimulateService(name string) {
	fmt.Printf("Service %s aborted\n", name)
}

func CommitChanges(success bool) string {
	if success {
		go CommitSimulateService("A")
		go CommitSimulateService("B")
		go CommitSimulateService("C")
		time.Sleep(1 * time.Second)
		return "Operation commited successfully!"
	} else {
		go AbortSimulateService("A")
		go AbortSimulateService("B")
		go AbortSimulateService("C")
		time.Sleep(1 * time.Second)
		return "Operation aborted successfully!"
	}
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
	fmt.Println("\n\nNew Request")

	test := req.GetTest()

	A := test[0] == '1'
	B := test[1] == '1'
	C := test[2] == '1'

	outA := make(chan bool)
	outB := make(chan bool)
	outC := make(chan bool)

	go func() { outA <- PrepareSimulateService("A", A) }()
	go func() { outB <- PrepareSimulateService("B", B) }()
	go func() { outC <- PrepareSimulateService("C", C) }()

	success := isSuccessful(<-outA, <-outB, <-outC)
	message := CommitChanges(success)

	return &pb.OperationResponse{Success: success, Message: message}, nil
}

func main() {
	server := grpc.NewServer()
	transactionManager := &TransactionManagerServer{}
	pb.RegisterTransactionManagerServer(server, transactionManager)

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	fmt.Println("Starting Transaction Manager server...")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
