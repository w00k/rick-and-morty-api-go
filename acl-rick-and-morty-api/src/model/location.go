package model

type Location struct {
	Id        int      `json:"id"`
	Name      string   `json:"name"`
	Type      string   `json:"type"`
	Dimension string   `json:"dimension"`
	Residents []string `json:"residents"`
	Url       string   `json:"url"`
	Created   string   `json:"created"`
}
