package main

import (
	"arman-task/internal/controller"
	"arman-task/internal/services/grpc"
	"fmt"
)

func main() {
	// create grpc client
	es := grpc.NewClient()

	// start user input
	fmt.Println("---- Add new user ----")
	for {
		controller.AddUserAndSegmention(es)
	}
}
