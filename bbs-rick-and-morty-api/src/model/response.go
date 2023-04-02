package model

type Response struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Status       string `json:"status"`
	Species      string `json:"species"`
	ResponseType string `json:"type"`
	EpisodeCount int    `json:"episode_count"`
	MyOrigin     Origin `json:"origin"`
}
