package model

import (
	"database/sql"
	"encoding/json"
	"github.com/lib/pq"
	"log"
)

func extractStatusData(rows *sql.Rows, status string) (map[string][]string, error) {
	dataMap := make(map[string][]string)
	defer rows.Close()
	for rows.Next() {
		var date string
		var availableByte []byte
		var unAvailableByte []string
		if err := rows.Scan(&date, &availableByte, pq.Array(&unAvailableByte)); err != nil {
			log.Printf("error while trying to extract data from the db rows: %v", err)
			continue
		}
		if status == "available" {
			var data map[string]float64
			if err := json.Unmarshal(availableByte, &data); err != nil {
				log.Printf("error while trying to unmarshal data: %v", err)
				continue
			}
			var extract []string
			for key := range data {
				extract = append(extract, key)
			}
			dataMap[date] = extract
			continue
		}
		dataMap[date] = unAvailableByte
	}
	return dataMap, nil
}

func extractTimeData(rows *sql.Rows, timeTaken string) (map[string][]string, error) {
	dataMap := make(map[string][]string)
	defer rows.Close()
	for rows.Next() {
		var date string
		var availableByte []byte
		if err := rows.Scan(&date, &availableByte); err != nil {
			log.Printf("error while trying to extract data from the db rows: %v", err)
			continue
		}
		var data map[string]float64
		if err := json.Unmarshal(availableByte, &data); err != nil {
			log.Printf("error while trying to unmarshal data: %v", err)
			continue
		}
		dataMap[date] = extractWithTimeFilter(data, timeTaken)
	}
	return dataMap, nil
}

func extractWithTimeFilter(data map[string]float64, filter string) []string {
	var extract []string
	for key, value := range data {
		if filter == "less" && value <= 1 {
			extract = append(extract, key)
			continue
		}
		if filter == "greater" && value > 1 {
			extract = append(extract, key)
			continue
		}
	}
	return extract
}
