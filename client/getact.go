/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"log"
	"main/activity_pb"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

// getactCmd represents the getact command
var getactCmd = &cobra.Command{
	Use:   "getact",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
		handleError(err)
		defer conn.Close()
		c := activity_pb.NewUserServiceClient(conn)
		email, err := cmd.Flags().GetString("email")
		if err != nil {
			log.Fatal(err)
		}
		GetActivity(c, email)
	},
}

func init() {
	rootCmd.AddCommand(getactCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	getactCmd.PersistentFlags().String("email", "", "User email")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getactCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
