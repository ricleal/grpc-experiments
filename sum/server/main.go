package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"exp/sum"
)

type server struct {
	sum.UnimplementedSumServer
}

func (s *server) Add(ctx context.Context, req *sum.SumRequest) (*sum.SumResponse, error) {
	if req.Num1 == 0 || req.Num2 == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "Both numbers cannot be zero")
	}
	return &sum.SumResponse{Result: req.Num1 + req.Num2}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	sum.RegisterSumServer(s, &server{})
	log.Println("Server listening on port 8080")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
