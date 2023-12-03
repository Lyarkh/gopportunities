package main

import (
	"github.com/Lyarkh/gopportunities/config"
	"github.com/Lyarkh/gopportunities/router"
)

func main() {
	// Initialize configs
	err := config.Init()
	if err != nil {
		panic(err)
	}

	// Initialize router
	router.Initialize()
}
