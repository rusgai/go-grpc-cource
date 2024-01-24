package domain

// is implantation type a Rocket
type Rocket struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Type    string `json:"type"`
	Flights int    `json:"flights"`
}

func NewRocket(id, name, type_ string, flights int) *Rocket {
	return &Rocket{id, name, type_, flights}
}
