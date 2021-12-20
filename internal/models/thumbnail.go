package models

// Thumbnail struct
type Thumbnail struct {
	BaseModel
	FileName       string `json:"file_name"`
	BucketName     string `json:"bucket_name"`
	ServerLocation string `json:"server_location"`
	FileSize       string `json:"file_size"`
	Room           Room   `json:"room"`
	RoomId         uint64 `json:"room_id"`
	Hotel          Hotel  `json:"hotel"`
	HotelId        uint64 `json:"hotel_id"`
}
