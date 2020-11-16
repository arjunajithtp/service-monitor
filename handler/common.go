package handler

import (
	"fmt"
	"github.com/arjunajithtp/service-monitor/model"
	"time"
)

func validateDate(date, tim string) bool {
	_, err := time.Parse(time.RFC3339, fmt.Sprintf("%sT%sZ", date, tim))
	if err != nil {
		return false
	}
	return true
}

// Connector holds the information required for contacting the service
type Connector struct {
	getByDate      func(string, string, string) (map[string][]string, error)
	getByTimeTaken func(string, string, string) (map[string][]string, error)
}

func getByDate(from, to, status string) (map[string][]string, error) {
	return model.GetByDate(from, to, status)
}

func getByTimeTaken(from, to, timeTaken string) (map[string][]string, error) {
	return model.GetByTimeTaken(from, to, timeTaken)
}

// GetConnector supplies a fresh connector
func GetConnector() *Connector {
	return &Connector{
		getByDate:      getByDate,
		getByTimeTaken: getByTimeTaken,
	}
}
