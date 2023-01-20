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

func check_string(t *testing.T, got string, want string) {
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func check_bool(t *testing.T, got bool, want bool) {
	if got != want {
		s := fmt.Sprint("got", got, ", wanted", want)
		t.Errorf(s)
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
	check_string(t, got, want)
	got = UserAdd(c, "testuser2", "testuser2@gmail.com", 13131313)
	want = "User already exist"
	check_string(t, got, want)
}

func TestActivityAdd(t *testing.T) {
	c, conn := connectToServer()
	defer conn.Close()

	got := ActivityAdd(c, "testuser1@gmail.com", "Sleep", 7, "label2")
	want := "User activity added"
	check_string(t, got, want)
}

func TestGetUser(t *testing.T) {
	c, conn := connectToServer()
	defer conn.Close()

	got := GetUser(c, "testuser1@gmail.com")
	want := true
	check_bool(t, got, want)
	got = GetUser(c, "unknownuser@gmail.com")
	want = false
	check_bool(t, got, want)

}

func TestGetActivity(t *testing.T) {
	c, conn := connectToServer()
	defer conn.Close()

	got := GetActivity(c, "sai@gmail.com")
	want := true

	check_bool(t, got, want)
}

func TestUpdateUser(t *testing.T) {
	c, conn := connectToServer()
	defer conn.Close()
	name := "hemanth sai 123"
	email := "h@gmail.com"
	phone := 909090909
	got := UpdateUser(c, email, name, int64(phone))
	want := "User details updated"
	check_string(t, got, want)
}

func TestRemoveUser(t *testing.T) {
	c, conn := connectToServer()
	defer conn.Close()

	got := RemoveUser(c, "testuser1@gmail.com")
	want := "User deleted successfully"
	check_string(t, got, want)
}
