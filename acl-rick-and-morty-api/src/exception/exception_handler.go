package exception

import (
	"io/ioutil"
	"log"
	"net/http"
	"w00k/rick-and-morty-api-go/acl-rick-and-morty-api/src/model"
)

func HttpResponseHandler(response *http.Response, errorInput error, endpointName string) (body []byte, exception model.Exception) {

	if errorInput != nil {
		log.Printf("%s - error to call Rick and Morty API", endpointName)
		exception.Status = 500
		exception.Message = errorInput.Error()
	}

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Printf("%s - %s", endpointName, err.Error())
		exception.Status = 409
		exception.Message = "Error with API response"
	}

	return
}
