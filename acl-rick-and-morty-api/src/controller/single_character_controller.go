package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	service "github.com/w00k/rick-and-morty-api-go/acl-rick-and-morty-api/src/service"
)

func GetSingleCharacter(c *gin.Context) {
	singleCharacter, err := service.CallSingleCharacter(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, singleCharacter)
		return
	}

	c.JSON(http.StatusOK, singleCharacter)
}
