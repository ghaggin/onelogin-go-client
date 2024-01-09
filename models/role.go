package models

type Role struct {
	ID     int    `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Admins []int  `json:"admins,omitempty"`
	Apps   []int  `json:"apps,omitempty"`
	Users  []int  `json:"users,omitempty"`
}
