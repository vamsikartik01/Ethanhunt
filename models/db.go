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
	RoomId     int    `json:"room_id"`
}

type Devices struct {
	Id         int    `json:"id"`
	Name       string `josn:"name"`
	Mode       string `json:"mode"`
	Status     string `json:"status"`
	HubId      int    `json:"hub_id"`
	RoomId     int    `json:"room_id"`
	AccountSid int    `json:"accountSid"`
	Value      string `json:"value"`
}
