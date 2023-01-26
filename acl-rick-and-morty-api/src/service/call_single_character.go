package service

import (
	"encoding/json"
	"net/http"
	exceptionHandler "w00k/rick-and-morty-api-go/acl-rick-and-morty-api/src/exception"
	"w00k/rick-and-morty-api-go/acl-rick-and-morty-api/src/model"
)

const endpointSingleCharacter string = "https://rickandmortyapi.com/api/character"

func CallSingleCharacter(characterId string) (data model.SingleCharacter, exception model.Exception) {
	response, err := http.Get(endpointSingleCharacter + "/" + characterId)

	body, exception := exceptionHandler.HttpResponseHandler(response, err, "CallSingleCharacter")

	defer response.Body.Close()

	if response.StatusCode != 200 {
		json.Unmarshal(body, &exception)
		exception.Status = response.StatusCode
		return
	}
	json.Unmarshal(body, &data)

	return
}
