package dto

type PaginationFilter struct {
	Page             int `json:"page"`
	PerPage          int `json:"per_page"`
	IgnorePagination bool
	TenantID         uint64
}

type Operator string

const (
	eq Operator = "eq"
	gt Operator = "gt"
	lt Operator = "lt"
)

type Filter struct {
	Field    interface{}
	Value    interface{}
	Operator Operator
}

type PaginationFilterInput struct {
	PaginationFilter
	Filters []*Filter
}

func (f PaginationFilter) ParseOperator(operator Operator) string {

	switch operator {
	case eq:
		return "="
	case gt:
		return ">="
	case lt:
		return "<="
	}

	return ""
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
