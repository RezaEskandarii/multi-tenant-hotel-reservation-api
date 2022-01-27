package models

import "os"

// Thumbnail struct
type Thumbnail struct {
	BaseModel
	FileName       string   `json:"file_name"`
	BucketName     string   `json:"bucket_name"`
	ServerLocation string   `json:"server_location"`
	Room           Room     `json:"room"  gorm:"foreignKey:RoomId;references:id"`
	RoomId         uint64   `json:"room_id"`
	Hotel          Hotel    `json:"hotel"  gorm:"foreignKey:HotelId;references:id"`
	HotelId        uint64   `json:"hotel_id"`
	VersionID      string   `json:"version_id" gorm:"type:varchar(255)"`
	FileSize       int64    `json:"file-size" gorm:"type:varchar(255)"`
	File           *os.File `json:"file" gorm:"-"`
}
