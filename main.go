package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"rest-api/controller"
)

func main() {
	r := gin.Default()
	r.GET("/", controller.Welcome)
	r.GET("/ping", controller.Ping)
	r.GET("/health", controller.Status)

	r.NoRoute(controller.Show404)
	if err := r.Run(); err != nil {
		log.Fatal("Looks like something went wrong ")
	}

}
