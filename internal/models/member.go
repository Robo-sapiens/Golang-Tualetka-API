package models

//easyjson:json
type Member struct {
	UserID   		 int    `json:"userID"`
	RoomID           int    `json:"roomID,omitempty"`
	Nickname 		 string `json:"nickname"`
	Name     	    *string `json:"name,omitempty"`
	Phone            string `json:"phone"`
	Status           string `json:"status"`
	ToiletPaperCount int    `json:"toiletPaperCount"`
	PayAbility 		 bool   `json:"payAbility"`
	Valuable 		 bool   `json:"valuable"`
}

//easyjson:json
type Paper struct {
	UserID     int `json:"userID"`
	RoomID     int `json:"roomID"`
	PaperCount int `json:"paperCount"`
}

//easyjson:json
type Valuable struct {
	UserID   int  `json:"userID"`
	RoomID   int  `json:"roomID"`
	Valuable bool `json:"valuable"`
}

//easyjson:json
type PayAbility struct {
	UserID     int  `json:"userID"`
	RoomID     int  `json:"roomID"`
	PayAbility bool `json:"payAbility"`
}

//easyjson:json
type Members []*Member