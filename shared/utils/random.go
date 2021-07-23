package utils

import (
	"math/rand"
	"time"
)

// GetRandomInt util function
func GetRandomInt(start int, end int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(end-start) + start
}
