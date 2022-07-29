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

func (r Response) SetId(id int) {
	r.id = id
}

func (r Response) GetId() int {
	return r.id
}

func (r Response) SetName(name string) {
	r.name = name
}

func (r Response) GetName() string {
	return r.name
}

func (r Response) SetStatus(status string) {
	r.status = status
}

func (r Response) GetStatus() string {
	return r.status
}

func (r Response) SetSpecies(species string) {
	r.species = species
}

func (r Response) GetSpecies() string {
	return r.species
}

func (r Response) SetResponseType(responseType string) {
	r.responseType = responseType
}

func (r Response) GetResponseType() string {
	return r.responseType
}

func (r Response) SetEpisode(episodeCount int) {
	r.episodeCount = episodeCount
}

func (r Response) GetEpisodeCount() int {
	return r.episodeCount
}

func (r Response) SetOrigin(origin Origin) {
	r.origin = origin
}

func (r Response) GetOrigin() Origin {
	return r.origin
}
