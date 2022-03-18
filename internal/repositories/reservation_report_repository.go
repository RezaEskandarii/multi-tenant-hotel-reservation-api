package repositories

import (
	"gorm.io/gorm"
	"reservation-api/internal/dto"
	"reservation-api/internal/models"
	"strings"
)

type ReportRepository struct {
	DB *gorm.DB
}

func NewReportRepository(db *gorm.DB) *ReportRepository {
	return &ReportRepository{DB: db}
}

func (r *ReportRepository) GetReservations(filter *dto.ReservationFilter) (error, *dto.ReservationReport) {

	reservations := make([]*models.Reservation, 0)

	query := r.DB.Model(&models.Reservation{})
	query = r.getReservationFilteredQuery(query, filter)
	query = paginateQuery(query, filter.Page, filter.PerPage)

	if err := query.Scan(&reservations).Error; err != nil {
		return err, nil
	}

	return nil, &dto.ReservationReport{
		Data:    reservations,
		Filters: filter,
	}

}

func (r *ReportRepository) getReservationFilteredQuery(query *gorm.DB, filter *dto.ReservationFilter) *gorm.DB {

	if filter.CreatedFrom != nil {
		query = query.Where("created_at >= ?", filter.CreatedFrom)
	}

	if filter.CreatedTo != nil {
		query = query.Where("created_at <= ?", filter.CreatedTo)
	}

	if filter.CheckInFrom != nil {
		query = query.Where("checkin_date >= ?", filter.CheckInFrom)
	}

	if filter.CheckInTo != nil {
		query = query.Where("checkin_date >= ?", filter.CheckInTo)
	}

	if strings.TrimSpace(filter.GuestName) != "" {
		query = query.Where("supervisor.first_name LIKE '%?%' OR "+
			"supervisor.middle_name LIKE '%?%' OR "+
			"supervisor.last_name LIKE '%?%'",
			filter.GuestName, filter.GuestName, filter.GuestName)
	}

	if filter.RateCodeId != 0 {
		query = query.Where("rate_code_id=?", filter.RateCodeId)
	}

	if filter.RoomId != 0 {
		query = query.Where("room_id=?", filter.RateCodeId)
	}

	if filter.CheckStatus != nil {
		query = query.Where("check_status=?", filter.CheckStatus)
	}

	return query
}
