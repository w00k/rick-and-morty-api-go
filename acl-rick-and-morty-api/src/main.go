package main

import (
	"github.com/gin-gonic/gin"
	controller "github.com/w00k/rick-and-morty-api-go/acl-rick-and-morty-api/src/controller"
)

func main() {
	r := gin.Default()
	r.GET("/single-character/:id", controller.GetSingleCharacter)
	r.Run(":8080")
}
