package models

type JamesInfo struct {
	AccountSid string `json:"account_sid"`
	HubId      string `json:"hub_id"`
	DeviceId   string `json:"device_id"`
	Status     string `json:"status"`
	Value      string `json:"value"`
}
