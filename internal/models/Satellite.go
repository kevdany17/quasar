package models

type SatelliteRequest struct {
	Name     string   `json:"name"`
	Distance float32  `json:"distance"`
	Message  []string `json:"message"`
}

type SatelliteResponse struct {
	Position Position `json:"position"`
	Message  string   `json:"message"`
}

type Request struct {
	Satellites []SatelliteRequest `json:"satellites"`
}

type Position struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
}

type Satellite struct {
	Name string  `json:"name"`
	X    float32 `json:"x"`
	Y    float32 `json:"y"`
}

type SatellitesList []Satellite

type Storage struct {
	Satellite SatelliteRequest `json:"satellite"`
	Position  Position         `json:"position"`
	Message   string           `json:"message"`
}

var DataBase []Storage

var LoadCarrier = Satellite{
	Name: "LoadCarrier",
	X:    100,
	Y:    75.5,
}

func NewSatellitesList() SatellitesList {
	list := make(SatellitesList, 3)

	list[0] = Satellite{
		Name: "Kenobi",
		X:    -500,
		Y:    -200,
	}

	list[1] = Satellite{
		Name: "Skywalker",
		X:    100,
		Y:    -100,
	}

	list[2] = Satellite{
		Name: "Sato",
		X:    500,
		Y:    100,
	}

	return list

}
