# User-Activity-Logger

User Activity Logger is a CLI application to to track users daily activities.Daily activities includes, “play”, “sleep”, “eat” and “read”. Each record will have activity type, time spent and timestamp of the activity creation.

## Tech Stack
- Golang
- gRPC
- Protobuf
- MongoDB
- Cobra-cli

## Features
- Create User
- Add Activity
- Get User Details
- Get Activity Details
- Delete User
- Update User
- Activity IsDone
- Activity IsValid

 ## Setup
  
  Installing protocol compiler
  
    $ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
    $ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
   
  Update PATH
  
    $ export PATH="$PATH:$(go env GOPATH)/bin"
  
  Installing grpc
  
    $ go get -u google.golang.org/grpc
  
  Installing Mongo-driver
    
    $ go get go.mongodb.org/mongo-driver/mongo
  
  Installing godotenv
     
    $ go get github.com/joho/godotenv
   
  Generating gRPC code
  
    $ protoc --go_out=. --go_opt=paths=source_relative \
      --go-grpc_out=. --go-grpc_opt=paths=source_relative \
      activity_pb/activity.proto
  
  Installing cobra-cli
  
    $ go install github.com/spf13/cobra-cli@latest
   

## Running the application
  
  ### Start server
  
     $ go run server/myserver.go
     
  ### Run client CLI
     
     $ cd client
     $ go install
     
  ### Client CLI commands
     
  #### To add user
     $ client adduser <name> <email> <phone-number>
  
  #### To add activity
     $ client addact <email> <activitytype> <duration> <label>
     
  #### To update activity
     $ client adduser <name> <email> <phone-number>
  
  #### To get user details
     $ client getuser --email=<email>
  
  #### To get user activity
     $ client getact --email=<email>
  
  #### To check if user activity isdone
     $ client isdone --email=<email> <activity-type>
     
  #### To check is user activity isvalid
     $ client isvalid --email=<email> <activity-type>
  
  #### To delete user
     $ client deluser --email=<email>
  
    
     
     
     
   


