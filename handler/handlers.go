package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/arjunajithtp/service-monitor/model"
	"github.com/kataras/iris/v12"
	"github.com/lib/pq"
	"log"
	"net/http"
)

// Handler is for handling the search requests
func Handler(ctx iris.Context) {
	fromDate := ctx.URLParam("fromDate")
	fromTime := ctx.URLParam("fromTime")
	from := fmt.Sprintf("%s %s", fromDate, fromTime)
	if !validateDate(fromDate, fromTime) {
		ctx.StatusCode(http.StatusBadRequest)
		ctx.WriteString("incorrect format for from date or time")
		return
	}
	toDate := ctx.URLParam("toDate")
	toTime := ctx.URLParam("toTime")
	to := fmt.Sprintf("%s %s", toDate, toTime)
	if !validateDate(toDate, toTime) {
		ctx.StatusCode(http.StatusBadRequest)
		ctx.WriteString("incorrect format for to date or time")
		return
	}

	status := ctx.URLParam("status")
	if status != "" && (status != "available" && status != "unavailable") {
		ctx.StatusCode(http.StatusBadRequest)
		ctx.WriteString("status can only be available or unavailable")
		return
	}
	rows, err := model.GetByDate(from, to)
	if err != nil {
		ctx.StatusCode(http.StatusNotFound)
		ctx.WriteString(err.Error())
		return
	}

	data, err := extractStatusData(rows, status)
	if err != nil {
		ctx.StatusCode(http.StatusInternalServerError)
		ctx.WriteString(err.Error())
		return
	}

	ctx.StatusCode(http.StatusOK)
	ctx.JSON(data)
}

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
