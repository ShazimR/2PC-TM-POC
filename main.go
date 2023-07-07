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

func CommitSimulateService(name string) bool {
	var delay int = rand.Intn(2)
	time.Sleep(time.Duration(delay) * time.Second)
	fmt.Printf("Service %s commited\n", name)
	return true // is completed
}

func AbortSimulateService(name string) bool {
	var delay int = rand.Intn(2)
	time.Sleep(time.Duration(delay) * time.Second)
	fmt.Printf("Service %s aborted\n", name)
	return true // is completed
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

	// Phase 1:
	go func() { outA <- PrepareSimulateService("A", A) }()
	go func() { outB <- PrepareSimulateService("B", B) }()
	go func() { outC <- PrepareSimulateService("C", C) }()

	success := isSuccessful(<-outA, <-outB, <-outC)

	// Phase 2:
	var message string
	if success {
		fmt.Println("\nStarting Commit")
		go func() { outA <- CommitSimulateService("A") }()
		go func() { outB <- CommitSimulateService("B") }()
		go func() { outC <- CommitSimulateService("C") }()
		message = "Operation commited successfully!"
	} else {
		fmt.Println("\nStarting Abort")
		go func() { outA <- AbortSimulateService("A") }()
		go func() { outB <- AbortSimulateService("B") }()
		go func() { outC <- AbortSimulateService("C") }()
		message = "Operation aborted successfully!"
	}
	<-outA
	<-outB
	<-outC
	fmt.Println("Operation Complete")

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
