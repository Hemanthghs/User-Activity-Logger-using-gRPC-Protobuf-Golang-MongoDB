package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type user_item struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone int64  `json:"phone"`
}

type activity_item struct {
	ActivityType string `json:"activitytype"`
	Duration     int32  `json:"duration"`
	Label        string `json:"label"`
	Timestamp    string `json:"timestamp"`
	Email        string `json:"email"`
}

type email_item struct {
	Email string `json:"email"`
}

type activity_check struct {
	Email        string `json:"email"`
	ActivityType string `json:"activitytype"`
}

var user_collection *mongo.Collection
var activity_collection *mongo.Collection

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func goDotEnvVariable(key string) string {
	err := godotenv.Load("../.env")
	handleError(err)
	return os.Getenv(key)
}

func initializeMigration() {
	godotenv.Load(".env")
	mongo_uri := goDotEnvVariable("MONGODB_URI")
	client, err := mongo.NewClient(options.Client().ApplyURI(mongo_uri))
	handleError(err)

	fmt.Println("MongoDB Connected")

	err = client.Connect(context.TODO())
	handleError(err)

	//selecting the user data collection
	user_collection = client.Database("useractivity").Collection("userdata")

	//selecting the activity data collection
	activity_collection = client.Database("useractivity").Collection("useractivitydata")
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user user_item
	json.NewDecoder(r.Body).Decode(&user)
	email := user.Email
	filter := bson.M{
		"email": email,
	}
	var result_data []user_item
	cursor, err := user_collection.Find(context.TODO(), filter)
	handleError(err)
	cursor.All(context.Background(), &result_data)
	if len(result_data) != 0 {
		w.Write([]byte("User already exist"))
		return
	}
	user_collection.InsertOne(context.TODO(), user)
	fmt.Println(user)
	w.Write([]byte("User added successfully"))
}

func AddActivity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var activity activity_item
	json.NewDecoder(r.Body).Decode(&activity)
	activity_collection.InsertOne(context.TODO(), activity)
	w.Write([]byte("User activity added successfully"))

}

func RemoveUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user_email email_item

	json.NewDecoder(r.Body).Decode(&user_email)
	filter := bson.M{
		"email": user_email.Email,
	}
	activity_collection.DeleteMany(context.TODO(), filter)
	u_r, err := user_collection.DeleteOne(context.TODO(), filter)
	handleError(err)
	if u_r.DeletedCount == 0 {
		w.Write([]byte("User does not exist"))
	} else {
		w.Write([]byte("User deleted successfully"))
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user_email email_item
	json.NewDecoder(r.Body).Decode(&user_email)
	filter := bson.M{
		"email": user_email.Email,
	}
	var result_data []user_item
	cursor, err := user_collection.Find(context.TODO(), filter)
	handleError(err)
	cursor.All(context.Background(), &result_data)
	if len(result_data) == 0 {
		w.Write([]byte("User does not exist"))
	} else {
		result := fmt.Sprint(" Name: ", result_data[0].Name, "\nEmail: ", result_data[0].Email, "\nPhone: ", result_data[0].Phone)
		w.Write([]byte(result))
	}

}

func GetActivity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user_email email_item
	json.NewDecoder(r.Body).Decode(&user_email)
	filter := bson.M{
		"email": user_email.Email,
	}
	var result_data []activity_item
	cursor, err := activity_collection.Find(context.TODO(), filter)
	handleError(err)
	cursor.All(context.Background(), &result_data)
	if len(result_data) == 0 {
		w.Write([]byte("User activities does not exist"))
	} else {
		for i := 0; i < len(result_data); i++ {
			result := fmt.Sprint("Email: ", result_data[i].Email, "\nActivityType: ", result_data[i].ActivityType, "\nTimestamp: ", result_data[i].Timestamp, "\nDuration: ", result_data[i].Duration, "\nLabel: ", result_data[i].Label)
			w.Write([]byte("\n-------------------\n"))
			w.Write([]byte(result))
		}
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user user_item
	json.NewDecoder(r.Body).Decode(&user)
	email := user.Email
	filter := bson.M{
		"email": email,
	}
	update := bson.D{{"$set", bson.D{{"email", user.Email}, {"name", user.Name}, {"phone", user.Phone}}}}
	_, err := user_collection.UpdateOne(context.Background(), filter, update)

	handleError(err)
	result := "User details updated"

	w.Write([]byte(result))
}

func ActivityIsValid(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var act activity_check
	json.NewDecoder(r.Body).Decode(&act)
	filter := bson.M{
		"email":        act.Email,
		"activitytype": act.ActivityType,
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
	w.Write([]byte(result))
}

func ActivityIsDone(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var act activity_check
	json.NewDecoder(r.Body).Decode(&act)
	filter := bson.M{
		"email":        act.Email,
		"activitytype": act.ActivityType,
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
	w.Write([]byte(result))
}

func initializeRouter() {
	r := mux.NewRouter()
	r.HandleFunc("/adduser", AddUser).Methods("POST")
	r.HandleFunc("/addact", AddActivity).Methods("POST")
	r.HandleFunc("/deluser", RemoveUser).Methods("POST")
	r.HandleFunc("/getuser", GetUser).Methods("POST")
	r.HandleFunc("/getact", GetActivity).Methods("POST")
	r.HandleFunc("/updateuser", UpdateUser).Methods("POST")
	r.HandleFunc("/actisvalid", ActivityIsValid).Methods("POST")
	r.HandleFunc("/actisdone", ActivityIsDone).Methods("POST")

	log.Fatal(http.ListenAndServe("0.0.0.0:9000", r))
}

func main() {
	initializeMigration()
	initializeRouter()
}
