package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	service "github.com/w00k/rick-and-morty-api-go/acl-rick-and-morty-api/src/service"
)

func GetSingleCharacter(c *gin.Context) {
	singleCharacter, exception := service.CallSingleCharacter(c.Param("id"))
	if exception.Status != 0 {
		c.JSON(http.StatusBadRequest, exception)
		return
	}
	c.JSON(http.StatusOK, singleCharacter)
}
