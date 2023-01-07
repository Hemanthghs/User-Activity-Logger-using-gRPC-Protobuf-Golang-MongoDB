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

func pushUserToDb(ctx context.Context, item user_item) primitive.ObjectID {
	res, err := collection.InsertOne(ctx, item)
	handleError(err)

	return res.InsertedID.(primitive.ObjectID)
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
	docid := pushUserToDb(ctx, newUserItem)
	result := fmt.Sprintf("User: %v is created with docid %v", name, docid)

	userAddResponse := activity_pb.UserResponse{
		Result: result,
	}
	return &userAddResponse, nil
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

	collection = client.Database("useractivity").Collection("useractivity")

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch

	fmt.Println("Closing mongodb connection")
	if err := client.Disconnect(context.TODO()); err != nil {
		handleError(err)
	}

	s.Stop()

}
