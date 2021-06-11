package models

// Country country struct
type Country struct {
	baseModel
	Name  string `json:"name"`
	Alias string `json:"alias"`
}
