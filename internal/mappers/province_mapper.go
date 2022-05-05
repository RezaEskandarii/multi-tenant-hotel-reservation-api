package mappers

import (
	"reservation-api/internal/dto"
	"reservation-api/internal/models"
)

func (m Mapper) MapToCountry(dto dto.CountryDto) models.Country {
	result := models.Country{
		Name:  dto.Name,
		Alias: dto.Alias,
	}

	//if dto.Provinces != nil {
	//	for _, p := range dto.Provinces {
	//		province := &models.Province{
	//			Name:      p.Name,
	//			Alias:     p.Alias,
	//			Cities:    nil,
	//			CountryId: 0,
	//			Country:   nil,
	//		}
	//
	//		if p.Cities != nil {
	//			for _, c := range p.Cities {
	//				province.Cities = append(province.Cities, models.City{
	//					Name:       "",
	//					Alias:      "",
	//					ProvinceId: 0,
	//					Province:   nil,
	//				})
	//			}
	//		}
	//
	//		result.Provinces = append(result.Provinces, province)
	//	}
	//}

	return result
}
