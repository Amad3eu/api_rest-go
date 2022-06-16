package model

type Repository struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Visibility string `json:"visibility"`
	Star       int    `json:"star"`
}
