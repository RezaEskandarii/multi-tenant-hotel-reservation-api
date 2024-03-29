package commons

import "math"

var (
	defaultSize uint = 20
)

// PaginatedResult paginate list
type PaginatedResult struct {
	Records      interface{} `json:"records"`
	Page         uint        `json:"page"`
	PerPage      uint        `json:"per_page"`
	TotalRecords uint        `json:"total_records"`
	TotalPages   uint        `json:"total_pages"`
	Filters      interface{} `json:"filters"`
}

// NewPaginatedList it returns new paginatesList struct and fills fields.
func NewPaginatedList(totalTableRows uint, page uint, perPage uint) *PaginatedResult {

	p := &PaginatedResult{}

	if perPage == 0 {
		perPage = defaultSize
	}
	p.TotalRecords = totalTableRows
	//get upper bound page size
	p.TotalPages = uint(math.Ceil(float64(totalTableRows) / float64(perPage)))
	p.Page = page

	//fetch size , This determines how many lines to fetch from multi_tenancy_database
	p.PerPage = perPage

	//it sets offset value per every page by given HTTP GET  page parameter
	if p.Page > 1 {
		p.Page = (p.Page * perPage) - perPage
	}

	return p

}
