package domain

type Characters []Character

type Character struct {
	ID int `json:"id"`
	Name string `json:"name"`
}