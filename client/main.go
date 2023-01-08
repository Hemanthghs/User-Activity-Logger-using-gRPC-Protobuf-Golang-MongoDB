package main

import (
	"context"
	"fmt"
	"log"
	"main/activity_pb"
	"time"

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

func getTimeStamp() string {
	t := time.Now()
	ts := t.Format("01-02-2006 15:04:05 Monday")
	return ts
}

func ActivityAdd(c activity_pb.UserServiceClient) {
	t := time.Now()
	ts := t.Format("01-02-2006 15:04:05 Monday")
	activityAddRequest := activity_pb.ActivityRequest{
		Activity: &activity_pb.Activity{
			ActivityType: "Sleep",
			Timestamp:    ts,
			Duration:     1,
			Label:        "label1",
			Email:        "hemanth@gmail.com",
		},
	}

	res, err := c.ActivityAdd(context.Background(), &activityAddRequest)
	handleError(err)
	fmt.Println(res)
}

func ActivityIsValid(c activity_pb.UserServiceClient) {
	activityIsValidResquest := activity_pb.ActivityIsValidRequest{
		Email:        "hemanth@gmail.com",
		Activitytype: "Sleep",
	}
	res, err := c.ActivityIsValid(context.Background(), &activityIsValidResquest)
	handleError(err)
	fmt.Println(res)
}

func main() {
	fmt.Println(getTimeStamp())
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	handleError(err)
	fmt.Println("Client started")
	defer conn.Close()

	c := activity_pb.NewUserServiceClient(conn)
	// UserAdd(c)
	// ActivityAdd(c)
	ActivityIsValid(c)
}
