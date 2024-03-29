package controller

import (
	"log"
	"net/http"

	"w00k/rick-and-morty-api-go/acl-rick-and-morty-api/src/service"

	"github.com/gin-gonic/gin"
)

func GetSingleLocation(c *gin.Context) {
	singleLocation, exception := service.CallSingleLocation(c.Param("id"))
	log.Println("controller code: ", exception.Status)
	if exception.Status != 0 {
		c.JSON(http.StatusBadRequest, exception)
		return
	}
	c.JSON(http.StatusOK, singleLocation)
}
