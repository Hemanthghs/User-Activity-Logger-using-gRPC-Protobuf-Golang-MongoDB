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
			Email: "hemanth1@gmail.com",
			Phone: 1234567890,
		},
	}
	res, err := c.UserAdd(context.Background(), &userAddRequest)
	handleError(err)
	fmt.Println(res)
}

func ActivityAdd(c activity_pb.UserServiceClient) {
	activityAddRequest := activity_pb.ActivityRequest{
		Activity: &activity_pb.Activity{
			ActivityType: "Play",
			Timestamp:    "20:28 PM IST Jan 8 2023",
			Duration:     4,
			Label:        "label1",
			Email:        "hemanth@gmail.com",
		},
	}

	res, err := c.ActivityAdd(context.Background(), &activityAddRequest)
	handleError(err)
	fmt.Println(res)
}

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	handleError(err)
	fmt.Println("Client started")
	defer conn.Close()

	c := activity_pb.NewUserServiceClient(conn)
	// UserAdd(c)
	ActivityAdd(c)
}
