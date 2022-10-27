package structs

type weather struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}
type Data struct {
	Status weather `json:"status"`
}
