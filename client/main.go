package main

import (
	"client/services"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	creds := insecure.NewCredentials()

	cc, err := grpc.Dial(":50051", grpc.WithTransportCredentials(creds))
	if err != nil {
		panic(err)
	}
	defer cc.Close()
	startTime := time.Now()
	usersClient := services.NewUserServiceClient(cc)
	usersServer := services.NewUserServer(usersClient)

	data, err := usersServer.GetUser(&services.GetUserRequest{
		UserId: "test_1234",
	})
	if err != nil {
		panic(err)
	}
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	log.Printf("Time taken: %v", duration)
	println(data.UserId)

}
