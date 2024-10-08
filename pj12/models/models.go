package models

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Item struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	UserID int    `json:"user_id"`
}
