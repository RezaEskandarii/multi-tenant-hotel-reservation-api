package domain_services

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	mock_repositories "reservation-api"
	"reservation-api/internal/models"
	"testing"
)

func TestCountryService(t *testing.T) {

	testCases := []struct {
		name  string
		alias string
	}{
		{name: "iran", alias: "IR"},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repository := mock_repositories.NewMockCountryRepository(ctrl)
	countryService := NewCountryService(repository)

	for _, testCase := range testCases {
		result, err := countryService.Create(&models.Country{Name: testCase.name, Alias: testCase.alias})
		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, testCase.name, result.Name)
		assert.Equal(t, testCase.alias, result.Alias)
	}
}
