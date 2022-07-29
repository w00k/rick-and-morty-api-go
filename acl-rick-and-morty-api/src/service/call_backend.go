package service

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/w00k/rick-and-morty-api-go/acl-rick-and-morty-api/src/model"
)

const endpointSingleCharacter string = "https://rickandmortyapi.com/documentation/#get-a-single-character"

func CallSingleCharacter(characterId int) (data model.SingleCharacter, er error) {
	//var data model.SingleCharacter
	response, err := http.Get(endpointSingleCharacter + "/" + strconv.Itoa(characterId))

	//call api
	if err != nil {
		log.Println("CallSingleCharacter - error to call rick and Morty Api Single Character")
		er = err
	}
	defer response.Body.Close()

	//map json
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("CallSingleCharacter - %s", err.Error())
		er = err
	}
	json.Unmarshal(body, &data)
	log.Println("CallSingleCharacter - body : ", data)

	return
}
