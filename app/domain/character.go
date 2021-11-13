package domain

type Characters []Character

type Character struct {
	ID int `csv:"id"`
	Name string `csv:"name"`
}