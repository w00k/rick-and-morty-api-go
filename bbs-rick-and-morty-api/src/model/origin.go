package model

type Origin struct {
	name      string   `json:"name"`
	url       string   `json:"url"`
	dimension string   `json:"dimension"`
	resident  []string `json:"resident"`
}
