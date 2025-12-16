package handler

import (
	"chatapp/internal/controller"
	"chatapp/internal/model"
	pb "chatapp/pkg"
	"context"
	"log"
)

// EmployeeServer implements pb.EmployeeOpsServer
type EmployeeOpsServer struct {
	pb.UnimplementedEmployeeOpsServer
	Ctrler controller.Controller
}

func (s *EmployeeOpsServer) Add(ctx context.Context, req *pb.AddEmployeeRequest) (*pb.AddEmployeeResponse, error) {
	log.Println("handler ...")
	log.Printf("\nhandler .. req [%v]", req.Emp)
	if req.Emp == nil {
		return &pb.AddEmployeeResponse{Ok: false}, nil
	}

	emp := model.Employee{
		Name: req.Emp.Name,
		Age:  req.Emp.Age,
	}

	_ = s.Ctrler.Add(ctx, emp)
	return &pb.AddEmployeeResponse{Ok: true}, nil
}
