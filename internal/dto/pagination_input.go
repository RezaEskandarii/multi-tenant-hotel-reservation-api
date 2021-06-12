package dto

type PaginationInput struct {
	Page    int `json:"page"`
	PerPage int `json:"per_page"`
}
