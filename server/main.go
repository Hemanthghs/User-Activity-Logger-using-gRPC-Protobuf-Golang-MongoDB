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
	Id    primitive.ObjectID `bson:"_id.omitempty"`
	Name  string             `bson:"name"`
	Email string             `bson:"email"`
	Phone int64              `bson:"phone"`
}

type activity_item struct {
	Id    primitive.ObjectID `bson:"_id.omitempty"`
	Email string             `bson:"email"`
	// ActivityType int32              `bson:"activity_type"`
	Timestamp string `bson:"tiemstamp"`
	Duration  int32  `bson:"duration"`
	Label     string `bson:"label"`
}

func pushUserToDb(ctx context.Context, item user_item) string {
	email := item.Email
	filter := bson.M{
		"email": email,
	}

	var result_data []user_item
	cursor, err := collection.Find(context.TODO(), filter)
	handleError(err)

	cursor.All(context.Background(), &result_data)

	if len(result_data) != 0 {
		result := "User already exist"
		return result
	}

	collection.InsertOne(ctx, item)
	result := "User created"
	return result
}

func pushActivityToDb(ctx context.Context, item activity_item) string {
	_, err := collection.InsertOne(ctx, item)
	handleError(err)

	return "pushtoactivity called..."
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

func (*server) UserActivityAdd(ctx context.Context, req *activity_pb.ActivityRequest) (*activity_pb.ActivityResponse, error) {
	fmt.Println(req)
	// activity_type := req.GetActivity().GetActivityType()
	timestamp := req.GetActivity().GetTimestamp()
	duration := req.GetActivity().GetDuration()
	label := req.GetActivity().GetLabel()
	email := req.GetActivity().GetEmail()
	newAtivityItem := activity_item{
		Email: email,
		// ActivityType: activity_type,
		Timestamp: timestamp,
		Duration:  duration,
	}
	dbres := pushActivityToDb(ctx, newAtivityItem)
	result := fmt.Sprintf("--- added %v ----", dbres)
	userActivityAddResponse := activity_pb.ActivityResponse{
		Result: result,
	}
	// fmt.Println(activity_type)
	fmt.Println(timestamp)
	fmt.Println(duration)
	fmt.Println(label)
	fmt.Println(email)

	return &userActivityAddResponse, nil

}

func goDotEnvVariable(key string) string {
	err := godotenv.Load(".env")

	handleError(err)

	return os.Getenv(key)
}

var collection *mongo.Collection

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

	collection = client.Database("useractivity").Collection("activitystore")

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch

	fmt.Println("Closing mongodb connection")
	if err := client.Disconnect(context.TODO()); err != nil {
		handleError(err)
	}

	s.Stop()

}
