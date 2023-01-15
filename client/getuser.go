/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"main/activity_pb"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

// getuserCmd represents the getuser command
var getuserCmd = &cobra.Command{
	Use:   "getuser",
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
		email, err2 := cmd.Flags().GetString("email")
		handleError(err2)
		GetUser(c, email)
	},
}

func init() {
	rootCmd.AddCommand(getuserCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	getuserCmd.PersistentFlags().String("email", "", "User email")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getuserCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
