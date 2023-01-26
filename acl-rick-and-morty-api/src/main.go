package main

import (
	controller "w00k/rick-and-morty-api-go/acl-rick-and-morty-api/src/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/single-character/:id", controller.GetSingleCharacter)
	r.GET("/single-location/:id", controller.GetSingleLocation)

	r.Run(":8080")
}
