package model

type Origin struct {
	name      string
	url       string
	dimension string
	resident  []string
}

func (o Origin) SetName(name string) {
	o.name = name
}

func (o Origin) GetName() string {
	return o.name
}

func (o Origin) SetUrl(url string) {
	o.url = url
}

func (o Origin) GetUrl() string {
	return o.url
}

func (o Origin) SetDimension(dimension string) {
	o.dimension = dimension
}

func (o Origin) GetDimension() string {
	return o.dimension
}

func (o Origin) SetResident(resident []string) {
	o.resident = resident
}

func (o Origin) GetResident() []string {
	return o.resident
}

func (o Origin) Origin(name string, url string, dimension string, resident []string) {
	o.name = name
	o.url = url
	o.dimension = dimension
	o.resident = resident
}

func (o Origin) GetOrigin() {
	return origin
}
