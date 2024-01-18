package models

type Rooms struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	AccountSid int    `json:"accountSid"`
}

type Hubs struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	AccountSid int    `json:"accountSid"`
	RefId      string `json:"ref_id"`
}

type Devices struct {
	Id         int    `json:"id"`
	Name       string `josn:"name"`
	Mode       string `json:"mode"`
	Status     string `json:"status"`
	HubPort    int    `json:hub_port`
	HubId      int    `json:"hub_id"`
	RoomId     int    `json:"room_id"`
	AccountSid int    `json:"accountSid"`
	Value      bool   `json:"value"`
	IsFavorite bool   `json:"is_favorite"`
	Type       string `json:"type"`
}
