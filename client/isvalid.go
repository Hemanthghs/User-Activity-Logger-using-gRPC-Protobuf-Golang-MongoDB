/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"main/activity_pb"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

// isvalidCmd represents the isvalid command
var isvalidCmd = &cobra.Command{
	Use:   "isvalid",
	Short: "To check whether the user activity is valid or not.",
	Long: `To checkout whether the user activiti is valid or not.
	
Inputs:
	email, activity-type

Example:
	client isvalid --email=<email> <activity-type>
	`,
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
		handleError(err)
		defer conn.Close()
		c := activity_pb.NewUserServiceClient(conn)
		email, err2 := cmd.Flags().GetString("email")
		handleError(err2)
		activityType := args[0]
		ActivityIsValid(c, email, activityType)
	},
}

func init() {
	rootCmd.AddCommand(isvalidCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	isvalidCmd.PersistentFlags().String("email", "", "User email")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// isvalidCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
