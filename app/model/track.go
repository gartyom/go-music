package model

type Track struct {
	Title    string `json:"title"`
	Features string `json:"features"`
	Audio    string `json:"audio"`
	Duration string `json:"duration"`
	Number   string `jsono:"number"`
}
