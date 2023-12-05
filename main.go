package main

import (
	"fmt"

	"github.com/Lyarkh/gopportunities/config"
	"github.com/Lyarkh/gopportunities/router"
)

func main() {
	// Initialize configs
	err := config.Init()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Initialize router
	router.Initialize()
}
