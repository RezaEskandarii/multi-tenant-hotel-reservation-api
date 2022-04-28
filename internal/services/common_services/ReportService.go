package common_services

import (
	"errors"
	"fmt"
	"github.com/xuri/excelize/v2"
	"reflect"
	"reservation-api/internal/utils"
	"reservation-api/pkg/translator"
	"strings"
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

	err, slice := utils.ConvertToInterfaceSlice(input)
	if err != nil {
		return err, nil
	}

	f := excelize.NewFile()
	sheetName := "Sheet1"
	index := f.NewSheet(sheetName)
	rowIdx := 1

	item1 := slice[0]
	for i := 0; i < reflect.TypeOf(item1).NumField(); i++ {
		colName := fmt.Sprintf("%s%d", getColName(i), rowIdx)
		f.SetCellValue(sheetName, colName, reflect.TypeOf(item1).Field(i).Name)
	}

	for i := 0; i < itemsValue.Len(); i++ {
		item := reflect.Indirect(itemsValue.Index(i))
		if item.Kind() == reflect.Struct {
			row := reflect.Indirect(item)
			rowIdx++
			for j := 0; j < row.NumField(); j++ {
				var value any
				if row.CanInterface() {
					value = row.Field(j).Interface()
				} else {
					value = row.Field(j)
				}
				colName := fmt.Sprintf("%s%d", getColName(j), rowIdx)
				if value == nil || strings.Contains(fmt.Sprintf("%s", value), "<nil>") {
					value = ""
				}
				f.SetCellValue(sheetName, colName, value)
			}
		}
	}

	f.SetActiveSheet(index)
	// Save spreadsheet by the given path.
	if err := f.SaveAs("Book1.xlsx"); err != nil {
		fmt.Println(err)
	}

	return nil, nil
}

func getColName(i int) string {

	str := fmt.Sprintf("%d", i)
	strResult := strings.Builder{}
	for _, chr := range str {
		char := fmt.Sprintf("%c", chr)
		number, _ := utils.ConvertToUint(char)

		strResult.Write([]byte(fmt.Sprintf("%c", rune('A'-1+number))))
	}

	return strResult.String()
}
