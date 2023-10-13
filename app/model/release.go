package model

type Release struct {
	Uuid  string `json:"uuid"`
	Title string `json:"name"`
	Image string `json:"cover"`
}
