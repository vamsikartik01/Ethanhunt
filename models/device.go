package models

type DeviceInfo struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Mode   string `json:"mode"`
	Status string `json:"status"`
	Value  string `json:"value"`
}
