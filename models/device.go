package models

type DeviceInfo struct {
	Id      string `json:"id"`
	HubId   string `json:"hub_id"`
	HubPort string `json:"hub_port"`
	Value   string `json:"value"`
	Type    string `json:"type"`
	RefId   string `json:"ref_id"`
	Ping    string `json:"ping"`
}
