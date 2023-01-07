package main

import (
	"context"
	"fmt"
	"log"
	"main/activity_pb"

	"google.golang.org/grpc"
)

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func UserAdd(c activity_pb.UserServiceClient) {
	userAddRequest := activity_pb.UserRequest{
		User: &activity_pb.User{
			Name:  "hemanth",
			Email: "hemanth3@gmail.com",
			Phone: 1234567890,
		},
	}
	res, err := c.UserAdd(context.Background(), &userAddRequest)
	handleError(err)
	fmt.Println(res)
}

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	handleError(err)
	fmt.Println("Client started")
	defer conn.Close()

	c := activity_pb.NewUserServiceClient(conn)
	UserAdd(c)
}
