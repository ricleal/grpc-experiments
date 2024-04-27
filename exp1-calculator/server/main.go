package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"exp/common"
)

type server struct {
	common.UnimplementedCalculatorServer
}

func (s *server) Add(ctx context.Context, req *common.AddRequest) (*common.AddReply, error) {
	return &common.AddReply{N1: req.N1 + req.N2}, nil
}

func (s *server) Subtract(ctx context.Context, req *common.SubtractRequest) (*common.SubtractReply, error) {
	return &common.SubtractReply{N1: req.N1 - req.N2}, nil
}

func (s *server) Multiply(ctx context.Context, req *common.MultiplyRequest) (*common.MultiplyReply, error) {
	return &common.MultiplyReply{N1: req.N1 * req.N2}, nil
}

func (s *server) Divide(ctx context.Context, req *common.DivideRequest) (*common.DivideReply, error) {
	if req.N2 == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "cannot divide by zero")
	}
	return &common.DivideReply{N1: req.N1 / req.N2}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	common.RegisterCalculatorServer(s, &server{})
	log.Println("Server listening on port 8080")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
