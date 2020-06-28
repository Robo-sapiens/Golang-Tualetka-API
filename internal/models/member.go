package models

//easyjson:json
type Member struct {
	UserID   		 int `json:"userID"`
	Nickname 		 string `json:"nickname"`
	Name     	    *string `json:"name,omitempty"`
	Phone            string `json:"phone"`
	Status           string `json:"status"`
	ToiletPaperCount int `json:"toiletPaperCount"`
	PayAbility 		 bool `json:"payAbility"`
	Valuable 		 bool `json:"valuable"`
}