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
