package service

import (
	"encoding/json"
	"net/http"
	exceptionHandler "w00k/rick-and-morty-api-go/acl-rick-and-morty-api/src/exception"
	"w00k/rick-and-morty-api-go/acl-rick-and-morty-api/src/model"
)

const endpointSingleLocation string = "https://rickandmortyapi.com/api/location"

func CallSingleLocation(locationId string) (data model.Location, exception model.Exception) {
	response, err := http.Get(endpointSingleLocation + "/" + locationId)

	body, exception := exceptionHandler.HttpResponseHandler(response, err, "CallSingleLocation")

	defer response.Body.Close()

	if response.StatusCode != 200 {
		json.Unmarshal(body, &exception)
		exception.Status = response.StatusCode
		return
	}

	json.Unmarshal(body, &data)
	return
}
