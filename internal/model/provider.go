package model

type Provider struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Endpoint string `json:"endpoint"`
}
