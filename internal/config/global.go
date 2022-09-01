package config

import "time"

var (
	RoomDefaultLockMinute   float64 = 20
	RoomDefaultLockDuration         = time.Now().Add(time.Minute * 20)
	HotelsBucketName                = "hotels-bucket"
	EmailQueueName                  = "email_queue"
	ReservationQueueName            = "reservation_queue"
	SendEmailRetryCount     uint    = 3
	TenantID                        = "TenantID"
)
