package handler

import (
	pb "chatapp/pkg"
	"context"
	"log"
)

// CalculatorServer implements pb.CalculatorServer
type CalculatorServer struct {
	pb.UnimplementedCalculatorServer
}

func (s *CalculatorServer) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddResponse, error) {
	log.Printf("Finding %d+%d", req.Num1, req.Num2)
	return &pb.AddResponse{Sum: req.Num1 + req.Num2}, nil
}
