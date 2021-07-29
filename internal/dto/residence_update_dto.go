package dto

type ResidenceUpdateDto struct {
	Name             string  `json:"name" valid:"required"`
	PhoneNumber1     string  `json:"phone_number1" valid:"required"`
	PhoneNumber2     string  `json:"phone_number2"`
	ProvinceId       uint64  `json:"province_id" valid:"required"`
	CityId           uint64  `json:"city_id" valid:"required"`
	Address          string  `json:"address" valid:"required"`
	PostalCode       string  `json:"postal_code" valid:"required"`
	Longitude        float64 `json:"longitude" valid:"required"`
	Latitude         float64 `json:"latitude" valid:"required"`
	FaxNumber        string  `json:"fax_number"`
	Website          string  `json:"website"`
	EmailAddress     string  `json:"email_address" valid:"email"`
	Description      string  `json:"description"`
	ResidenceTypeId  uint64  `json:"residence_type_id" valid:"required"`
	ResidenceGradeId uint64  `json:"residence_grade_id" valid:"required"`
}
