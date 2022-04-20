package main

import (
	"arman-task/internal/controller"
	"arman-task/internal/services/grpc"
	"fmt"
)

func main() {
	es := grpc.NewClient()

	fmt.Println("---- Add new user ----")
	for {
		controller.AddUserAndSegmention(es)
	}
}
