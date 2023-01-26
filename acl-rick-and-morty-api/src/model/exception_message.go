package model

type Exception struct {
	Status  int    `json:"status"`
	Message string `json:"error"`
}
