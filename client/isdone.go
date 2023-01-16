/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"main/activity_pb"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

// isdoneCmd represents the isdone command
var isdoneCmd = &cobra.Command{
	Use:   "isdone",
	Short: "To check whether the user activity is done or not.",
	Long: `To checkout whether the user activiti is done or not.
	
Inputs:
	email, activity-type

Example:
	client isdone --email=<email> <activity-type>
	`,
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
		handleError(err)
		defer conn.Close()
		c := activity_pb.NewUserServiceClient(conn)
		email, err2 := cmd.Flags().GetString("email")
		handleError(err2)
		activityType := args[0]
		ActivityIsDone(c, email, activityType)
	},
}

func init() {
	rootCmd.AddCommand(isdoneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	isdoneCmd.PersistentFlags().String("email", "", "User email")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// isdoneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
