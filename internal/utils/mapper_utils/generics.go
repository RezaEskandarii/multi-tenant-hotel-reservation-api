package mapper_utils

import "encoding/json"

type Number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 |
		uint32 | uint64 | uintptr | float32 | float64
}

func ConvertByGeneric[T any](obj T, data []byte) T {
	json.Unmarshal(data, &obj)
	return obj
}
