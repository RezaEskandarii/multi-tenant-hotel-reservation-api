// Package global_variables contains variables to set and read configs, s
// set data and read data from Context and etx.

package global_variables

import "time"

var (
	RoomDefaultLockMinute   float64 = 20
	RoomDefaultLockDuration         = time.Now().Add(time.Minute * 20)
	HotelsBucketName                = "hotels-bucket"
	EmailQueueName                  = "email_queue"
	ReservationQueueName            = "reservation_queue"
	SendEmailRetryCount     uint    = 3
	TenantIDKey                     = "TenantID"
	TenantIDCtx                     = "TenantIDCtx"
	ClaimsKey                       = "Claims"
	CurrentLang                     = "CurrentLang"
	UserClaims                      = "user_claims"
	DefaultTenantID                 = uint64(1)
)
