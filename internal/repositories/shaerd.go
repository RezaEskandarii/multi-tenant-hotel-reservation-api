package repositories

import (
	"gorm.io/gorm"
	"math"
	"reflect"
	"reservation-api/internal/commons"
	"reservation-api/internal/dto"
)

// paginatedList apply pagination filters query amd return PaginatedResult
func paginatedList(model interface{}, db *gorm.DB, input *dto.PaginationFilter) (*commons.PaginatedResult, error) {

	var total int64

	modelType := reflect.TypeOf(model)
	functionResult := reflect.MakeSlice(reflect.SliceOf(modelType), 0, int(total))

	query := db.Model(model)
	query.Count(&total)

	modelSlice := reflect.New(functionResult.Type())
	modelSlice.Elem().Set(functionResult)

	result := commons.NewPaginatedList(uint(total), uint(input.Page), uint(input.PageSize))
	query = query.Limit(int(result.PerPage)).Offset(int(result.Page)).Order("id desc").Find(modelSlice.Interface())

	if query.Error != nil {

		return nil, query.Error
	}

	result.Records = modelSlice.Interface()
	return result, nil
}

func paginateWithFilter(query *gorm.DB, result interface{}, filters interface{}, pageNumber, pageSize int, ignorePagination bool) *commons.PaginatedResult {

	var count int64 = 0
	query.Count(&count)

	if ignorePagination == false {
		query = paginateQuery(query, pageSize, pageNumber).Scan(&result)
	}

	return &commons.PaginatedResult{
		Records:      result,
		Page:         uint(pageNumber),
		PerPage:      uint(pageSize),
		TotalRecords: uint(count),
		TotalPages:   uint(math.Ceil(float64(count) / float64(pageSize))),
		Filters:      filters,
	}
}

func paginateQuery(db *gorm.DB, page, pageSize int) *gorm.DB {
	offset := (page - 1) * pageSize
	return db.Offset(offset).Limit(pageSize)
}
