package main

import (
	"arman-task/internals/controller"
	"fmt"
)

func main() {
	fmt.Println("---- Add new user ----")
	for {
		controller.GetUserAndSegmention()
	}
}
