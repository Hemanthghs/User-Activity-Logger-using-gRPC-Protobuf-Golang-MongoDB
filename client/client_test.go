package main

import (
	"fmt"
	"main/activity_pb"
	"testing"

	"google.golang.org/grpc"
)

// var c activity_pb.UserServiceClient

// func connectToServer() {

// }

func gotWant(t *testing.T, got string, want string) {
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func connectToServer() (activity_pb.UserServiceClient, *grpc.ClientConn) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	handleError(err)
	// defer conn.Close()
	c := activity_pb.NewUserServiceClient(conn)
	return c, conn
}

func TestUserAdd(t *testing.T) {
	c, conn := connectToServer()
	defer conn.Close()

	got := UserAdd(c, "testuser1", "testuser1@gmail.com", 1212121212)
	want := "User already exist"
	gotWant(t, got, want)
	got = UserAdd(c, "testuser2", "testuser2@gmail.com", 13131313)
	want = "User already exist"
	gotWant(t, got, want)
}

func TestActivityAdd(t *testing.T) {
	c, conn := connectToServer()
	defer conn.Close()

	got := ActivityAdd(c, "testuser1@gmail.com", "Sleep", 7, "label2")
	want := "User activity added"
	gotWant(t, got, want)
}

func TestGetUser(t *testing.T) {
	c, conn := connectToServer()
	defer conn.Close()

	got := GetUser(c, "testuser1@gmail.com")
	want := true
	if got != want {
		s := fmt.Sprint("got", got, ", wanted", want)
		t.Errorf(s)
	}
	got = GetUser(c, "unknownuser@gmail.com")
	want = false
	if got != want {
		s := fmt.Sprint("got", got, ", wanted", want)
		t.Errorf(s)
	}

}
