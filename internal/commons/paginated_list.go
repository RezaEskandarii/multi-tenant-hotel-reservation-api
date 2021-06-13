package commons

import "math"

var (
	defaultSize uint = 20
)

// PaginatedList paginate list
type PaginatedList struct {
	Data         interface{} `json:"data"`
	Page         uint        `json:"page"`
	PerPage      uint        `json:"per_page"`
	TotalRecords uint        `json:"total_records"`
	TotalPages   uint        `json:"total_pages"`
}

// Paginate it returns new paginatesList struct and fills fields.
func (p *PaginatedList) Paginate(totalTableRows uint, perPage uint, page uint) *PaginatedList {

	if perPage == 0 {
		perPage = defaultSize
	}
	p.TotalRecords = totalTableRows
	//get upper bound page size
	p.TotalPages = uint(math.Ceil(float64(totalTableRows) / float64(perPage)))
	p.Page = page

	//fetch size , This determines how many lines to fetch from database
	p.PerPage = perPage
	//fetch size offset default value
	p.Page = 0

	//it sets offset value per every page by given HTTP GET  page parameter
	if p.Page > 1 {
		p.Page = (p.Page * perPage) - perPage
	}

	return p

}
