package main

import (
	"github.com/HarlonGarcia/golang-jobs/config"
	"github.com/HarlonGarcia/golang-jobs/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	err := config.Init()

	if err != nil {
		logger.Errorf("Error initializing the application: %v", err)
		return
	}

	router.Init()
}
