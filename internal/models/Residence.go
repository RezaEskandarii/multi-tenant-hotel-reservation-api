package models

type Residence struct {
	BaseModel
	Name         string   `json:"name"`
	PhoneNumber1 string   `json:"phone_number1"`
	PhoneNumber2 string   `json:"phone_number2"`
	Province     Province `json:"province"`
	ProvinceId   uint64   `json:"province_id" gorm:"foreignKey:Province"`
	Address      string   `json:"address"`
	PostalCode   string   `json:"postal_code"`
	Longitude    float64  `json:"longitude"`
	Latitude     float64  `json:"latitude"`
	FaxNumber    string   `json:"fax_number"`
	Website      string   `json:"website"`
	EmailAddress string   `json:"email_address"`
	Owner        User     `json:"owner"`
	OwnerId      uint64   `json:"owner_id" gorm:"foreignKey:Owner"`
}
