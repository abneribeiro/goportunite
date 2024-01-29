package main

import (

	"github.com/abneribeiro/goportunite/config"
	"github.com/abneribeiro/goportunite/router"
)

var (
	logger config.Logger
)

func main() {
	logger = *config.GetLogger("main")
	// initialize the config
	err := config.Init()
	if err != nil {
		logger.Errof("Config initializing error: %v", err)
		return
	}
	// Initialize the router
	router.Initialize()

}
