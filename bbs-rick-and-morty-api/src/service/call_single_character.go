package service

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/w00k/rick-and-morty-api-go/bbs-rick-and-morty-api/src/exception"
	"github.com/w00k/rick-and-morty-api-go/bbs-rick-and-morty-api/src/model"
)

const endpointSingleCharacter string = "http://localhost:8080/single-character"

func CallSingleCharacter(characterId string) (data model.SingleCharacter, exception exception.Exception) {
	response, err := http.Get(endpointSingleCharacter + "/" + characterId)

	if err != nil {
		log.Println("CallSingleCharacter - error to call rick and Morty Api Single Character")
		exception.Status = 500
		exception.Message = err.Error()
	} else {
		exception.Status = response.StatusCode
		exception.Message = "OK"
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("CallSingleCharacter - %s", err.Error())
	}

	if response.StatusCode != 200 {
		json.Unmarshal(body, &exception)
		return
	}

	json.Unmarshal(body, &data)

	return
}
