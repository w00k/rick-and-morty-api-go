package service

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/w00k/rick-and-morty-api-go/acl-rick-and-morty-api/src/exception"
	"github.com/w00k/rick-and-morty-api-go/acl-rick-and-morty-api/src/model"
)

const endpointSingleLocation string = "https://rickandmortyapi.com/api/location"

func CallSingleLocation(locationId string) (data model.Location, exception exception.Exception) {
	response, err := http.Get(endpointSingleLocation + "/" + locationId)
	if err != nil {
		log.Println("CallSingleLocation - error to call Rick and Morty Api Single Location")
		exception.Status = 500
		exception.Message = err.Error()
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("CallSingleCharacter - error transform data %s", err.Error())
	}

	if response.StatusCode != 200 {
		json.Unmarshal(body, &exception)
		exception.Status = response.StatusCode
		return
	}

	json.Unmarshal(body, &data)
	return
}
