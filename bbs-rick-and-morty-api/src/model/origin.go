package model

type Origin struct {
	Name      string   `json:"name"`
	Url       string   `json:"url"`
	Dimension string   `json:"dimension"`
	Resident  []string `json:"residents"`
}
