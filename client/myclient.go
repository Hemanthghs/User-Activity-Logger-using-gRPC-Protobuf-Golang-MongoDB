package main

import (
	"context"
	"fmt"
	"log"
	"main/activity_pb"
	"time"
)

/*
function to handle runtime errors

Input:

	error
*/
func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

/*
function to create add-user request

Inputs:

	name, email, phone-number
*/
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

/*
function to get currect timestamp

Returns:

	string (timestamp)
*/
func getTimeStamp() string {
	t := time.Now()
	ts := t.Format("01-02-2006 15:04:05 Monday")
	return ts
}

/*
function to create add-activity reqeust

Input:

	email, activitytype, duration, label
*/
func ActivityAdd(c activity_pb.UserServiceClient, email string, at string, duration int32, label string) {
	t := time.Now()
	ts := t.Format("01-02-2006 15:04:05 Monday")
	activityAddRequest := activity_pb.ActivityRequest{
		Activity: &activity_pb.Activity{
			Activitytype: at,
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

/*
function to create activity-is-valid request

Inputs:

	email, activitytype (Play, Sleep, Eat, Study)
*/
func ActivityIsValid(c activity_pb.UserServiceClient, email string, activitytype string) {
	activityIsValidRequest := activity_pb.ActivityIsValidRequest{
		Email:        email,
		Activitytype: activitytype,
	}
	res, err := c.ActivityIsValid(context.Background(), &activityIsValidRequest)
	handleError(err)
	fmt.Println(res)
}

/*
function to create activity-is-done request

Inputs:

	email, activitytype (Play, Sleep, Eat, Study)
*/
func ActivityIsDone(c activity_pb.UserServiceClient, email string, activitytype string) {
	activityIsDoneRequest := activity_pb.ActivityIsDoneRequest{
		Email:        email,
		Activitytype: activitytype,
	}
	res, err := c.ActivityIsDone(context.Background(), &activityIsDoneRequest)
	handleError(err)
	fmt.Println(res)
}

/*
function to create update-user request

Inputs:

	email, name, phone-number
*/
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

/*
function to create get-activity request

Inputs:

	email
*/
func GetActivity(c activity_pb.UserServiceClient, email string) {
	getActivityRequest := activity_pb.GetActivityRequest{
		Email: email,
	}
	res, err := c.GetActivity(context.Background(), &getActivityRequest)
	handleError(err)
	fmt.Println(res)
}

/*
function to create get-user-details request

Inputs:

	email
*/
func GetUser(c activity_pb.UserServiceClient, email string) {
	getUserRequest := activity_pb.GetUserRequest{
		Email: email,
	}
	res, err := c.GetUser(context.Background(), &getUserRequest)
	handleError(err)
	fmt.Println(res)
}

/*
function to create remove-user request

Inputs:

	email
*/
func RemoveUser(c activity_pb.UserServiceClient, email string) {
	removeUserRequest := activity_pb.RemoveUserRequest{
		Email: email,
	}
	res, err := c.RemoveUser(context.Background(), &removeUserRequest)
	handleError(err)
	fmt.Println(res)
}
