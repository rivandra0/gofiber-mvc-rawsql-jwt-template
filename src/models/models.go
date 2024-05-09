package models

type User struct {
	Id               string `json:"Id"`
	Email            string `json:"Email"`
	UserName         string `json:"Username"`
	Password         string `json:"Password"`
	GroupId          string `json:"GroupId"`
	Tier             string `json:"Tier"`
	IsEmailConfirmed bool   `json:"IsEmailConfirmed"`
	IsActive         bool   `json:"IsActive"`
	CreatedAt        string `json:"CreatedAt"`
	CreatedBy        string `json:"CreatedBy"`
	UpdatedAt        string `json:"UpdatedAt"`
	UpdatedBy        string `json:"UpdatedBy"`
}
