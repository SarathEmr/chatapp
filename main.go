package main

import (
	"chatapp/internal/controller"
	"chatapp/internal/handler"
	"chatapp/internal/model"
	"chatapp/internal/repo"
	pb "chatapp/pkg"

	"context"
	"fmt"
	"log"
	"net"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

// func handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "Hello, World!")
// }

func main() {

	time.Sleep(10 * time.Second)

	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://admin:root@mongodb:27017/?authSource=admin") // "admin" is the default db name

	// Connect to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("mongodb connect failed: ", err)
	}
	defer client.Disconnect(ctx)

	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("mongodb ping failed: ", err)
	}

	fmt.Println("Connected to MongoDB!")

	r := repo.New(*client)
	ctrler := controller.Controller{Repo: r}

	r.Insert(ctx, model.Employee{
		Name:        "Tom",
		Designation: "Engineer",
		Age:         22,
	})

	/*
		http.HandleFunc("/", handler)
		fmt.Println("Server is listening on port 3000...")
		if err := http.ListenAndServe(":3000", nil); err != nil {
			fmt.Println("Error starting server:", err)
		}
	*/

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterCalculatorServer(grpcServer, &handler.CalculatorServer{})
	pb.RegisterEmployeeOpsServer(grpcServer, &handler.EmployeeOpsServer{Ctrler: ctrler})
	log.Println("gRPC server listening on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
