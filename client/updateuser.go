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

// updateuserCmd represents the updateuser command
var updateuserCmd = &cobra.Command{
	Use:   "updateuser",
	Short: "To update existing user details.",
	Long: `To update existing user details.

Inputs:
	email, name, phone-number

Example:
	client updateuser --email=<email> <name> <phone-number>
	`,
	Run: func(cmd *cobra.Command, args []string) {
		email, err := cmd.Flags().GetString("email")
		if err != nil {
			log.Fatal(err)
		}
		phone, err := strconv.ParseInt(args[1], 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
		handleError(err)
		defer conn.Close()

		c := activity_pb.NewUserServiceClient(conn)

		UpdateUser(c, email, args[0], phone)
	},
}

func init() {
	rootCmd.AddCommand(updateuserCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	updateuserCmd.PersistentFlags().String("email", "", "User email")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateuserCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
