package services

import (
	"github.com/stretchr/testify/assert"
	"hotel-reservation/internal/models"
	"hotel-reservation/internal/repositories"
	"hotel-reservation/pkg/database"
	"testing"
)

var (
	db, _ = database.GetDb()
)

func TestCountryService(t *testing.T) {

	countryService := *NewCountryService()
	countryService.Repository = *repositories.NewCountryRepository(db)

	t.Run("Create", func(t *testing.T) {
		c := &models.Country{
			Name:  "Iran",
			Alias: "IRI",
		}

		result, err := countryService.Create(c)

		assert.Nil(t, err)
		assert.Equal(t, result.Name, c.Name)
	})
}
