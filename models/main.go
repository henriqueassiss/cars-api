package models

type Car struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Brand string `json:"brand"`
}

var Cars []Car
