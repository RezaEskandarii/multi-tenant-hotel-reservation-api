package common_services

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"github.com/xuri/excelize/v2"
	"reflect"
	"reservation-api/internal/utils/mapper_utils"
	"strconv"
	"strings"
)

type ReportService struct {
}

func NewReportService() *ReportService {
	return &ReportService{}
}

// ExportToExcel gives a slice of structs
// this method's output is byte array and error
func (r *ReportService) ExportToExcel(input interface{}, lang string) ([]byte, error) {

	itemsValue := reflect.ValueOf(input)
	// return error if input type is struct
	if itemsValue.Kind() != reflect.Slice {
		return nil, errors.New("input is not type of interface")
	}

	// check input length
	if itemsValue.Len() == 0 {
		return nil, errors.New("input data length is 0")
	}

	// check if input contains none struct item
	for i := 0; i < itemsValue.Len(); i++ {
		if itemsValue.Index(i).Kind() != reflect.Struct {
			return nil, errors.New(fmt.Sprintf("the number of item {%d} type is not struct", i))
		}
	}

	// convert input to slice of structs
	slice, err := mapper_utils.ConvertToInterfaceSlice(input)
	if err != nil {
		return nil, err
	}

	f := excelize.NewFile()
	defer f.Close()

	sheetName := "Sheet1"
	index := f.NewSheet(sheetName)
	rowIdx := 1

	// get index 0 of slice to read fields of  struct and put field names as a Excel output's header.
	item1 := slice[0]

	for i := 0; i < reflect.TypeOf(item1).NumField(); i++ {
		// excel output headers col name
		colName := fmt.Sprintf("%s%d", r.getColName(i), rowIdx)
		f.SetCellValue(sheetName, colName, reflect.TypeOf(item1).Field(i).Name)
	}

	// get each item of given input
	for i := 0; i < itemsValue.Len(); i++ {

		item := reflect.Indirect(itemsValue.Index(i))
		if item.Kind() == reflect.Struct {

			row := reflect.Indirect(item)
			rowIdx++

			for j := 0; j < row.NumField(); j++ {
				// put field value into value field
				var value any

				if row.CanInterface() {
					value = row.Field(j).Interface()
				} else {
					value = row.Field(j)
				}
				// get excel column column name to put data
				colName := fmt.Sprintf("%s%d", r.getColName(j), rowIdx)

				if value == nil || strings.Contains(fmt.Sprintf("%s", value), "<nil>") {
					value = ""
				}

				f.SetCellValue(sheetName, colName, value)
			}
		}
	}

	f.SetActiveSheet(index)

	buffer, err := f.WriteToBuffer()

	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil

}

func (*ReportService) ExportToPDF(data interface{}, headers []string) ([]byte, error) {
	pdf := gofpdf.New("P", "mm", "A4", "")

	pdf.SetHeaderFunc(func() {
		pdf.SetFont("Arial", "B", 16)
		pdf.Cell(0, 10, "Report")
		pdf.Ln(20)
		pdf.SetFont("Arial", "B", 12)
		for _, header := range headers {
			pdf.Cell(40, 10, header)
		}
		pdf.Ln(12)
	})

	// Table column widths
	//colWidths := []float64{20, 50, 50, 20}
	//
	//// Draw table headers
	//for i, header := range headers {
	////	pdf.CellFormat(colWidths[i], 10, header, "1", 0, "", false, 0, "")
	//}

	pdf.SetFooterFunc(func() {
		pdf.SetY(-15)
		pdf.SetFont("Arial", "I", 8)
		pdf.Cell(0, 10, fmt.Sprintf("Page %d", pdf.PageNo()))
	})

	slice := reflect.ValueOf(data)
	pdf.AddPage()

	for row := 0; row < slice.Len(); row++ {
		//	pdf.AddPage()
		//	pdf.SetFont("Arial", "", 12)
		rowData := slice.Index(row)
		for col := 0; col < rowData.NumField(); col++ {
			value := rowData.Field(col).Interface()
			if value != nil && value != "" && fmt.Sprintf("%v", rowData) != value {
				pdf.Cell(40, 10, fmt.Sprintf("%v", value))
			}
		}
		pdf.Ln(8)
	}

	buf := new(bytes.Buffer)
	err := pdf.Output(buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// getColName returns excel column name per given column number
// For example, if input is 1, output will be A
// or if input is 12, output will be AB
// or if input is 2, output will be B
// or if input is 11, output will be AA
func (r *ReportService) getColName(i int) string {

	str := fmt.Sprintf("%d", i)
	strResult := strings.Builder{}

	for _, chr := range str {
		char := fmt.Sprintf("%c", chr)
		number, err := strconv.ParseInt(char, 10, 64)
		if err != nil {
			return ""
		}
		// generate output result
		strResult.Write([]byte(fmt.Sprintf("%c", rune('A'-1+number))))
	}

	return strResult.String()
}
