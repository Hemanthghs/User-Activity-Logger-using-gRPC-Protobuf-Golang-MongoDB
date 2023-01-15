package main

import (
	"context"
	"fmt"
	"log"
	"main/activity_pb"
	"net"
	"os"
	"os/signal"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type server struct {
	activity_pb.UnimplementedUserServiceServer
}

type user_item struct {
	Id    primitive.ObjectID `bson:"_id,omitempty"`
	Name  string             `bson:"name"`
	Email string             `bson:"email"`
	Phone int64              `bson:"phone"`
}

type activity_item struct {
	Id           primitive.ObjectID `bson:"_id,omitempty"`
	ActivityType string             `bson:"activity_type"`
	Duration     int32              `bson:"duration"`
	Label        string             `bson:"label"`
	Timestamp    string             `bson:"timestamp"`
	Email        string             `bson:"email"`
}

func pushUserToDb(ctx context.Context, item user_item) string {
	email := item.Email
	filter := bson.M{
		"email": email,
	}

	var result_data []user_item
	cursor, err := user_collection.Find(context.TODO(), filter)
	handleError(err)

	cursor.All(context.Background(), &result_data)

	if len(result_data) != 0 {
		result := "User already exist"
		return result
	}
	user_collection.InsertOne(ctx, item)
	result := "User created"
	return result
}

func pushActivityToDb(ctx context.Context, item activity_item) string {
	activity_collection.InsertOne(ctx, item)
	result := "User activity added"
	return result
}

func (*server) UserAdd(ctx context.Context, req *activity_pb.UserRequest) (*activity_pb.UserResponse, error) {
	fmt.Println(req)
	name := req.GetUser().GetName()
	email := req.GetUser().GetEmail()
	phone := req.GetUser().GetPhone()

	newUserItem := user_item{
		Name:  name,
		Email: email,
		Phone: phone,
	}
	dbres := pushUserToDb(ctx, newUserItem)
	result := fmt.Sprintf("%v", dbres)

	userAddResponse := activity_pb.UserResponse{
		Result: result,
	}
	return &userAddResponse, nil
}

func (*server) ActivityAdd(ctx context.Context, req *activity_pb.ActivityRequest) (*activity_pb.ActivityResponse, error) {
	fmt.Println(req)
	activity_type := req.GetActivity().GetActivityType()
	duration := req.GetActivity().GetDuration()
	label := req.GetActivity().GetLabel()
	timestamp := req.GetActivity().GetTimestamp()
	email := req.GetActivity().GetEmail()
	newActivityItem := activity_item{
		ActivityType: activity_type,
		Duration:     duration,
		Label:        label,
		Timestamp:    timestamp,
		Email:        email,
	}
	dbres := pushActivityToDb(ctx, newActivityItem)
	result := fmt.Sprintf("%v", dbres)

	activityAddResponse := activity_pb.ActivityResponse{
		Result: result,
	}

	return &activityAddResponse, nil

}

func (*server) ActivityIsValid(ctx context.Context, req *activity_pb.ActivityIsValidRequest) (*activity_pb.ActivityIsValidResponse, error) {
	fmt.Println(req)
	email := req.GetEmail()
	activity_type := req.GetActivitytype()
	filter := bson.M{
		"email":         email,
		"activity_type": activity_type,
	}
	var result_data []activity_item
	cursor, err := activity_collection.Find(context.Background(), filter)
	handleError(err)
	cursor.All(context.Background(), &result_data)
	var result string
	if len(result_data) == 0 {
		result = "No user exist with the given email"
	} else {
		if result_data[0].Duration > 2 {
			result = "Activity is Valid"
		} else {
			result = "Activity is Not Valid"
		}
	}

	activityIsValidResponse := activity_pb.ActivityIsValidResponse{
		Result: result,
	}
	return &activityIsValidResponse, nil
}

func (*server) ActivityIsDone(ctx context.Context, req *activity_pb.ActivityIsDoneRequest) (*activity_pb.ActivityIsDoneResponse, error) {
	fmt.Println(req)
	email := req.GetEmail()
	activity_type := req.GetActivitytype()
	filter := bson.M{
		"email":         email,
		"activity_type": activity_type,
	}
	var result_data []activity_item
	cursor, err := activity_collection.Find(context.Background(), filter)
	handleError(err)
	cursor.All(context.Background(), &result_data)
	var result string
	if len(result_data) == 0 {
		result = "No user exist with the given email"
	} else {
		if result_data[0].Duration > 4 {
			result = "Activity is Done"
		} else {
			result = "Activity is Not Done"
		}
	}

	activityIsDoneResponse := activity_pb.ActivityIsDoneResponse{
		Result: result,
	}
	return &activityIsDoneResponse, nil

}

func (*server) UpdateUser(ctx context.Context, req *activity_pb.UpdateUserRequest) (*activity_pb.UpdateUserResponse, error) {
	fmt.Println(req)
	email := req.GetUser().GetEmail()
	name := req.GetUser().GetName()
	phone := req.GetUser().GetPhone()

	filter := bson.M{
		"email": email,
	}

	update := bson.D{{"$set", bson.D{{"email", email}, {"name", name}, {"phone", phone}}}}
	_, err := user_collection.UpdateOne(context.Background(), filter, update)

	handleError(err)
	result := "User details updated"

	updateUserResponse := activity_pb.UpdateUserResponse{
		Result: result,
	}
	return &updateUserResponse, nil
}

func goDotEnvVariable(key string) string {
	err := godotenv.Load(".env")
	handleError(err)
	return os.Getenv(key)
}

func (*server) GetActivity(ctx context.Context, req *activity_pb.GetActivityRequest) (*activity_pb.GetActivityResponse, error) {
	fmt.Println(req)
	email := req.GetEmail()
	filter := bson.M{
		"email": email,
	}
	var result_data []activity_item
	cursor, err := activity_collection.Find(context.TODO(), filter)
	handleError(err)
	cursor.All(context.Background(), &result_data)
	if len(result_data) == 0 {
		getActivityResponse := activity_pb.GetActivityResponse{
			Status:   false,
			Activity: nil,
		}
		return &getActivityResponse, nil
	} else {
		getActivityResponse := activity_pb.GetActivityResponse{
			Status: true,
			Activity: &activity_pb.Activity{
				ActivityType: result_data[0].ActivityType,
				Timestamp:    result_data[0].Timestamp,
				Duration:     result_data[0].Duration,
				Label:        result_data[0].Label,
				Email:        result_data[0].Email,
			},
		}
		return &getActivityResponse, nil
	}
}

func (*server) GetUser(ctx context.Context, req *activity_pb.GetUserRequest) (*activity_pb.GetUserResponse, error) {
	fmt.Println(req)
	email := req.GetEmail()
	filter := bson.M{
		"email": email,
	}
	var result_data []user_item
	cursor, err := user_collection.Find(context.TODO(), filter)
	handleError(err)
	cursor.All(context.Background(), &result_data)
	if len(result_data) == 0 {
		getUserResponse := activity_pb.GetUserResponse{
			Status: false,
			User:   nil,
		}
		return &getUserResponse, nil
	} else {
		getUserResponse := activity_pb.GetUserResponse{
			Status: true,
			User: &activity_pb.User{
				Email: result_data[0].Email,
				Name:  result_data[0].Name,
				Phone: result_data[0].Phone,
			},
		}
		return &getUserResponse, nil
	}
}

var user_collection *mongo.Collection
var activity_collection *mongo.Collection

func main() {
	godotenv.Load(".env")

	opts := []grpc.ServerOption{}
	s := grpc.NewServer(opts...)

	activity_pb.RegisterUserServiceServer(s, &server{})

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	handleError(err)

	go func() {
		if err := s.Serve(lis); err != nil {
			handleError(err)
		}
	}()

	mongo_uri := goDotEnvVariable("MONGODB_URI")
	client, err := mongo.NewClient(options.Client().ApplyURI(mongo_uri))
	handleError(err)

	fmt.Println("MongoDB Connected")

	err = client.Connect(context.TODO())
	handleError(err)

	user_collection = client.Database("useractivity").Collection("userdata")
	activity_collection = client.Database("useractivity").Collection("useractivitydata")

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch

	fmt.Println("Closing mongodb connection")
	if err := client.Disconnect(context.TODO()); err != nil {
		handleError(err)
	}

	s.Stop()

}
