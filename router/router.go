package router

import "github.com/gin-gonic/gin"

func Initialize() {
	router := gin.Default()
	initializeRoutes()
	router.Run(":8080")
}
