package dto

type PaginationFilter struct {
	Page    int `json:"page"`
	PerPage int `json:"per_page"`
}

func NewPaginatedInput(page int, perPage int) *PaginationFilter {

	if perPage <= 0 {
		perPage = 20
	}

	return &PaginationFilter{
		Page:    page,
		PerPage: perPage,
	}
}
