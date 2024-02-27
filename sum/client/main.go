package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"exp/sum"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect to server: %v", err)
	}
	defer conn.Close()

	client := sum.NewSumClient(conn)
	req := &sum.SumRequest{Num1: 5, Num2: 3}
	response, err := client.Add(context.Background(), req)
	if err != nil {
		log.Fatalf("could not add: %v", err)
	}

	fmt.Printf("Sum: %d\n", response.GetResult())
}
