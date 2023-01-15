package main

import (
	"context"
	"fmt"
	"log"
	"main/activity_pb"
	"time"
)

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func UserAdd(c activity_pb.UserServiceClient, name string, email string, phone int64) {
	userAddRequest := activity_pb.UserRequest{
		User: &activity_pb.User{
			Name:  name,
			Email: email,
			Phone: phone,
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

func ActivityAdd(c activity_pb.UserServiceClient, activityType string, duration int32, label string, email string) {
	t := time.Now()
	ts := t.Format("01-02-2006 15:04:05 Monday")
	activityAddRequest := activity_pb.ActivityRequest{
		Activity: &activity_pb.Activity{
			ActivityType: activityType,
			Timestamp:    ts,
			Duration:     duration,
			Label:        label,
			Email:        email,
		},
	}

	res, err := c.ActivityAdd(context.Background(), &activityAddRequest)
	handleError(err)
	fmt.Println(res)
}

func ActivityIsValid(c activity_pb.UserServiceClient, email string, activityType string) {
	activityIsValidRequest := activity_pb.ActivityIsValidRequest{
		Email:        email,
		Activitytype: activityType,
	}
	res, err := c.ActivityIsValid(context.Background(), &activityIsValidRequest)
	handleError(err)
	fmt.Println(res)
}

func ActivityIsDone(c activity_pb.UserServiceClient, email string, activityType string) {
	activityIsDoneRequest := activity_pb.ActivityIsDoneRequest{
		Email:        email,
		Activitytype: activityType,
	}
	res, err := c.ActivityIsDone(context.Background(), &activityIsDoneRequest)
	handleError(err)
	fmt.Println(res)
}

func UpdateUser(c activity_pb.UserServiceClient, email string, name string, phone int64) {
	updateUserRequest := activity_pb.UpdateUserRequest{
		User: &activity_pb.User{
			Name:  name,
			Email: email,
			Phone: phone,
		},
	}
	res, err := c.UpdateUser(context.Background(), &updateUserRequest)
	handleError(err)
	fmt.Println(res)

}

func GetActivity(c activity_pb.UserServiceClient, email string) {
	getActivityRequest := activity_pb.GetActivityRequest{
		Email: email,
	}
	res, err := c.GetActivity(context.Background(), &getActivityRequest)
	handleError(err)
	fmt.Println(res)
}

func GetUser(c activity_pb.UserServiceClient, email string) {
	getUserRequest := activity_pb.GetUserRequest{
		Email: email,
	}
	res, err := c.GetUser(context.Background(), &getUserRequest)
	handleError(err)
	fmt.Println(res)
}

func RemoveUser(c activity_pb.UserServiceClient, email string) {
	removeUserRequest := activity_pb.RemoveUserRequest{
		Email: email,
	}
	res, err := c.RemoveUser(context.Background(), &removeUserRequest)
	handleError(err)
	fmt.Println(res)
}
