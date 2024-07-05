package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"server/services"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

func main() {
	// ตั้งค่า MongoDB
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	// เชื่อมต่อกับ collection
	db := client.Database("testdb").Collection("user-grpc")

	// สร้าง gRPC Server และลงทะเบียน UserService
	s := grpc.NewServer()
	services.RegisterUserServiceServer(s, services.NewUserServer(db))

	// เริ่มต้น server
	fmt.Println("server start listening on :50051")
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
