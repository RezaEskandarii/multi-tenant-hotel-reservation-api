package services

import (
	"github.com/stretchr/testify/assert"
	"hotel-reservation/internal/dto"
	"hotel-reservation/internal/models"
	"hotel-reservation/internal/repositories"
	"hotel-reservation/pkg/database"
	"testing"
)

var (
	db, _      = database.GetDb(true)
	pagination = dto.PaginationInput{
		Page:    1,
		PerPage: 20,
	}
)

func TestCountryService(t *testing.T) {

	id := uint64(1000000001)
	countryService := *NewCountryService()
	countryService.Repository = repositories.NewCountryRepository(db)

	t.Run("Create", func(t *testing.T) {
		c := &models.Country{
			Name:  "Iran",
			Alias: "IRI",
		}
		c.Id = id
		result, err := countryService.Create(c)
		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, result.Name, c.Name)
	})

	t.Run("Find", func(t *testing.T) {
		result, err := countryService.Find(id)
		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, result.Id, id)
		assert.NotEqual(t, result.Id, 0)
	})

	t.Run("FindAll", func(t *testing.T) {
		result, err := countryService.FindAll(&pagination)
		assert.Nil(t, err)
		assert.NotNil(t, result)
	})
}
