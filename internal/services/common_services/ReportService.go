package common_services

import (
	"errors"
	"fmt"
	"reflect"
	"reservation-api/pkg/translator"
)

type ReportService struct {
	Translator translator.TranslateService
}

func NewReportService(translateService translator.TranslateService) *ReportService {
	return &ReportService{Translator: translateService}
}

func (r *ReportService) ExportToExcel(input interface{}, lang string) (error, []byte) {

	itemsValue := reflect.ValueOf(input)
	if itemsValue.Kind() != reflect.Slice {
		return errors.New("input is not type of interface"), nil
	}

	if itemsValue.Len() == 0 {
		return errors.New("input data length is 0"), nil
	}

	for i := 0; i < itemsValue.Len(); i++ {
		if itemsValue.Index(i).Kind() != reflect.Struct {
			return errors.New(fmt.Sprintf("the number of item {%d} type is not struct", i)), nil
		}
	}

	for i := 0; i < itemsValue.Len(); i++ {
		item := itemsValue.Index(i)
		if item.Kind() == reflect.Struct {
			row := reflect.Indirect(item)
			for j := 0; j < row.NumField(); j++ {

				if row.CanInterface() {
					fmt.Println(row.Field(j).Interface())
				} else {
					fmt.Println(row.Field(j))
				}
			}
		}
	}
	return nil, nil
}
