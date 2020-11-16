package handler

import (
	"fmt"
	"time"
)

func validateDate(date, tim string) bool {
	_, err := time.Parse(time.RFC3339, fmt.Sprintf("%sT%sZ", date, tim))
	if err != nil {
		return false
	}
	return true
}
