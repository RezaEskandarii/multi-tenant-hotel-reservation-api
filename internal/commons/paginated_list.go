package commons

type PaginatedList struct {
	Data         interface{} `json:"data"`
	Page         int         `json:"page"`
	PerPage      int         `json:"per_page"`
	TotalRecords int         `json:"total_records"`
	TotalPages   int         `json:"total_pages"`
}
