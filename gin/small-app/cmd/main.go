package main

import (
	"github.com/gin-gonic/gin"
	"small-app/handlers"
)

// /data?id=2

func main() {
	r := gin.New()
	r.Use(gin.Logger())
	r.GET("/home")
	r.GET("/data", handlers.GetUserGin)

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")

}
