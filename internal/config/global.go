package config

import "time"

var (
	RoomDefaultLockMinute   = float64(20)
	RoomDefaultLockDuration = time.Now().Add(time.Minute * 20)
)
