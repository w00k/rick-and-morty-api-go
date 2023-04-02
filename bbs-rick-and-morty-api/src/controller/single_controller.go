package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/w00k/rick-and-morty-api-go/bbs-rick-and-morty-api/src/model"
	"github.com/w00k/rick-and-morty-api-go/bbs-rick-and-morty-api/src/service"
)

func SingleController(c *gin.Context) {

	singleCharacter, exception := service.CallSingleCharacter(c.Param("id"))

	if exception.Status != 200 {
		c.JSON(http.StatusBadRequest, exception)
		return
	}

	locationId := getLocationId(singleCharacter.Location.Url)
	log.Println("Locarion Id: ", locationId)

	location, err := service.CallSingleLocation(locationId)

	if err.Status != 0 {
		log.Println("Location Service Error")
		log.Println(location)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	response := buildResponse(singleCharacter, location)

	c.JSON(http.StatusOK, response)
}

func getLocationId(url string) string {

	indexBase := 41 // https://rickandmortyapi.com/api/location/3 len 42 before 3
	endBase := len(url)

	log.Printf("index: %d \nend: %d \n", indexBase, endBase)
	if endBase > indexBase {
		return url[indexBase:endBase]
	}
	return "0"
}

func buildResponse(singleCharacter model.SingleCharacter, location model.Location) (response model.Response) {

	episodeCount := 0

	if singleCharacter.Episode != nil {
		episodeCount = len(singleCharacter.Episode)
	}

	response.Id = singleCharacter.Id
	response.Name = singleCharacter.Name
	response.Status = singleCharacter.Status
	response.Species = singleCharacter.Species
	response.ResponseType = singleCharacter.Type
	response.EpisodeCount = episodeCount
	response.MyOrigin.Name = singleCharacter.Origin.Name
	response.MyOrigin.Url = singleCharacter.Origin.Url
	response.MyOrigin.Dimension = location.Dimension
	response.MyOrigin.Resident = location.Residents

	return
}
