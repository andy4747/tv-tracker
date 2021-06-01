package util

import (
	"time"
)

func GetCurrentDate() string {
	currentTime := time.Now()
	date := currentTime.Format("2006-01-02 15:04:05")
	return date
}

func IsTimeExpired(unixTime int64) bool {
	return unixTime < time.Now().Unix()
}
