package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"rest-api/config"
	"rest-api/controller"
)

func main() {
	environment := flag.String("e", "development", "")
	config.Init(*environment)

	r := gin.Default()

	r.GET("/", controller.Welcome)
	r.GET("/ping", controller.Ping)
	r.GET("/health", controller.Status)

	r.NoRoute(controller.Show404)

	err:= r.Run(config.GetConfig().GetString("server_port"));if err!=nil{
		fmt.Println("Looks like something went wrong ")
	}

}
