package models

//easyjson:json
type User struct {
	ID 		 int `json:"id"`
	Nickname string `json:"nickname"`
	Name 	*string `json:"name,omitempty"`
	Phone 	 string `json:"phone"`
	Status   string `json:"status"`
}