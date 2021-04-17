package controller

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func Ping(c *gin.Context) {
	c.JSON(200, Response{Message: "It's up and running", Code: 200})
}
func Welcome(c *gin.Context) {
	c.JSON(200, Response{Message: "Hello! Welcome to api.trinadhkoya.in", Code: 200})
}
func Status(c *gin.Context) {
	c.JSON(200, Response{Message: "Health is OK!", Code: 200})
}

func Show404(c *gin.Context) {
	c.JSON(400, Response{Message: "Oops! Page is not available", Code: 404})
}
