package repositories

import (
	"gorm.io/gorm"
	"hotel-reservation/internal/commons"
	"hotel-reservation/internal/dto"
	"reflect"
)

//finAll
func finAll(model interface{}, db *gorm.DB, input *dto.PaginationInput) (*commons.PaginatedList, error) {

	var total int64

	modelType := reflect.TypeOf(model)
	functionResult := reflect.MakeSlice(reflect.SliceOf(modelType), 0, int(total))

	query := db.Model(model)
	query.Count(&total)

	modelSlice := reflect.New(functionResult.Type())
	modelSlice.Elem().Set(functionResult)

	result := commons.NewPaginatedList(uint(total), uint(input.Page), uint(input.PerPage))
	query = query.Limit(int(result.PerPage)).Offset(int(result.Page)).Order("id desc").Find(modelSlice.Interface())

	if query.Error != nil {

		return nil, query.Error
	}

	result.Records = modelSlice.Interface()
	return result, nil
}
