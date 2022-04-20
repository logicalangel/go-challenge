package controller

import (
	"arman-task/internal/services"
	"arman-task/internal/usecase"
	"fmt"
	"log"

	"github.com/manifoldco/promptui"
)

func AddUserAndSegmention(es services.EstimateService) {
	promptUsername := promptui.Prompt{
		Label: "Enter Username",
	}

	username, err := promptUsername.Run()
	if err != nil {
		log.Fatal("failed to input username: ", err.Error())
	}

	promptSegmention := promptui.Prompt{
		Label:   "Enter Segment",
		Default: "non",
	}

	segment, err := promptSegmention.Run()
	if err != nil {
		log.Fatal("failed to input segment: ", err.Error())
	}

	err = usecase.SendUserSegmention(es, username, segment)
	if err != nil {
		log.Fatal("failed to sent segment to es: ", err.Error())
	}

	fmt.Println("user added successfully\n------------------------------")
}
