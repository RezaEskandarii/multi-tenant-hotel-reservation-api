package dto

type PaginationInput struct {
	Page    int `json:"page"`
	PerPage int `json:"per_page"`
}

func NewPaginatedInput(page int, perPage int) *PaginationInput {

	if perPage <= 0 {
		perPage = 20
	}

	return &PaginationInput{
		Page:    page,
		PerPage: perPage,
	}
}
