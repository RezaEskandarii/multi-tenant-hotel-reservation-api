package dto

type PaginationFilter struct {
	Page             int `json:"page"`
	PageSize         int `json:"page_size"`
	IgnorePagination bool
}

type PaginationFilterInput struct {
	PaginationFilter
}

func NewPaginatedInput(page int, perPage int) *PaginationFilter {

	if perPage <= 0 {
		perPage = 20
	}

	return &PaginationFilter{
		Page:     page,
		PageSize: perPage,
	}
}
