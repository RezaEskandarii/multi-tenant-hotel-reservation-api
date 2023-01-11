package tenant_dsn_resolver

import "reservation-api/internal/models"

func GetEntities() []interface{} {
	return []interface{}{
		models.Country{},
		models.City{},
		models.Province{},
		models.Currency{},
		models.User{},
		models.Hotel{},
		models.Room{},
		models.RoomType{},
		models.Guest{},
		models.RateGroup{},
		models.RateCode{},
		models.HotelGrade{},
		models.HotelType{},
		models.ReservationRequest{},
		models.Reservation{},
		models.Audit{},
		models.RateCodeDetail{},
		models.RateCodeDetailPrice{},
		models.Sharer{},
		models.Thumbnail{},
	}
}
