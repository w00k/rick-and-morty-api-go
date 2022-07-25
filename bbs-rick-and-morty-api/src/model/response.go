package model

type Response struct {
	id           int
	name         string
	status       string
	species      string
	responseType string
	episodeCount int
	origin       Origin
}
