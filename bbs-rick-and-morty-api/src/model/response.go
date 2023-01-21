package model

type Response struct {
	id           int    `json:"id"`
	name         string `json:"name"`
	status       string `json:"status"`
	species      string `json:"species"`
	responseType string `json:"type"`
	episodeCount int    `json:"episodie_count"`
	origin       Origin `json:"origin"`
}
