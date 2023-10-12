package model

type Release struct {
	Title  string `json:"name"`
	Artist string `json:"artist"`
	Type   string `json:"type"`
	Cover  string `json:"cover"`
	Tracks []Track
}
