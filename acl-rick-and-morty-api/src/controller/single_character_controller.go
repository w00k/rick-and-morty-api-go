package controller

import (
	"net/http"

	service "w00k/rick-and-morty-api-go/acl-rick-and-morty-api/src/service"

	"github.com/gin-gonic/gin"
)

func GetSingleCharacter(c *gin.Context) {
	singleCharacter, exception := service.CallSingleCharacter(c.Param("id"))
	if exception.Status != 0 {
		c.JSON(http.StatusBadRequest, exception)
		return
	}
	c.JSON(http.StatusOK, singleCharacter)
}
