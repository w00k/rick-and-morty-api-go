package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/w00k/rick-and-morty-api-go/bbs-rick-and-morty-api/src/controller"
)

func main() {
	log.Println("Start BBS")
	r := gin.Default()
	r.GET("/rick-and-morty/single/:id", controller.SingleController)

	r.Run(":8081")

}
