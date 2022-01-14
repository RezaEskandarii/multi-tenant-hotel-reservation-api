package config

import "time"

var (
	RoomDefaultLockDuration = time.Now().Add(time.Minute * 20)
)
