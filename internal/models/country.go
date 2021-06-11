package models

type Country struct {
	baseModel
	Name  string `json:"name"`
	Alias string `json:"alias"`
}
