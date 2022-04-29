package common_services

import (
	"github.com/stretchr/testify/assert"
	"reservation-api/internal/models"
	"testing"
)

func TestCanGetExcelOutput(t *testing.T) {

	reportService := NewReportService(nil)

	columnTestCases := []struct {
		given  int
		wanted string
	}{
		{
			given: 1, wanted: "A",
		},
		{
			given: 2, wanted: "B",
		},
		{
			given: 3, wanted: "C",
		},
		{
			given: 111, wanted: "AAA",
		},
		{given: 112, wanted: "AAB"},
	}

	t.Run("test_can_get_excel_output_for_slice_of_users", func(t *testing.T) {
		users := []models.User{
			{
				FirstName: "Reza",
				LastName:  "Eskandari",
				Username:  "rezaeskandari___",
				Email:     "test@test.test",
				IsActive:  true,
			},
		}

		output, err := reportService.ExportToExcel(users, "")

		assert.Nil(t, err)
		assert.NotNil(t, output)
		assert.True(t, len(output) > 0)
	})

	t.Run("test_can_ignore_none_struct_fields", func(t *testing.T) {
		strSlice := []string{
			"123",
			"aaa",
		}

		output, err := reportService.ExportToExcel(strSlice, "")

		assert.NotNil(t, err)
		assert.Nil(t, output)
		assert.Nil(t, output)
	})

	t.Run("test_can_generate_correct_column_name", func(t *testing.T) {

		for _, testCase := range columnTestCases {

			column := getColName(testCase.given)
			assert.Equal(t, column, testCase.wanted)
		}
	})
}
