package models

//easyjson:json
type Room struct {
	ID      int `json:"id"`
	Name    string `json:"name"`
	About  *string `json:"about"`
	Members []Member `json:"members"`
}