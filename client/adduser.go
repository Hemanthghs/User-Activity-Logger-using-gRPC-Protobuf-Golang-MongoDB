/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	"main/activity_pb"
	"strconv"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

// adduserCmd represents the adduser command
var adduserCmd = &cobra.Command{
	Use:   "adduser",
	Short: "To create a new user.",
	Long: `To create a new user and insert into the database.

Inputs:
	name, email, phone-number

Example:
	client adduser <name> <email> <phone-number>`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println(getTimeStamp())
		conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
		handleError(err)
		defer conn.Close()
		c := activity_pb.NewUserServiceClient(conn)
		phone, err := strconv.ParseInt(args[2], 10, 64)
		if err != nil {
			fmt.Println(err)
		}
		UserAdd(c, args[0], args[1], phone)
	},
}

func init() {
	rootCmd.AddCommand(adduserCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// adduserCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// adduserCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
