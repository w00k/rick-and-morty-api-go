package model

type SingleCharacter struct {
	Id       int      `json:"id"`
	Name     string   `json:"name"`
	Status   string   `json:"status"`
	Species  string   `json:"species"`
	Type     string   `json:"type"`
	Gender   string   `json:"gender"`
	Origin   origin   `json:"origin"`
	Location location `json:"location"`
	Image    string   `json:"image"`
	Episodes []string `json:"episodes"`
	Url      string   `json:"url"`
	Created  string   `json:"created"`
}
