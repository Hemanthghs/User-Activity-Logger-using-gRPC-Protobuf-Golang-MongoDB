/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"log"
	"main/activity_pb"
	"strconv"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

// addactCmd represents the addact command
var addactCmd = &cobra.Command{
	Use:   "addact",
	Short: "To add an activity of user",
	Long: `To add a specific activity of a user.
Activity types : (Play, Eat, Sleep, Study).

Inputs:
	email, activitytype, duration, label
	
Example:
	client addact <email> <activitytype> <duration> <label>
	`,
	Run: func(cmd *cobra.Command, args []string) {

		conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
		handleError(err)
		defer conn.Close()
		c := activity_pb.NewUserServiceClient(conn)
		// ActivityAdd(c, "Eat", 4, "label4", "hemanth5@gmail.com")
		duration, err := strconv.ParseInt(args[1], 10, 32)
		if err != nil {
			log.Fatal(err)
		}
		ActivityAdd(c, args[0], int32(duration), args[2], args[3])
	},
}

func init() {
	rootCmd.AddCommand(addactCmd)
	// rootCmd.PersistentFlags().String("email", "", "email of user")
}
