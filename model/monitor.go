package model

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/arjunajithtp/service-monitor/config"
	"github.com/lib/pq"
	"time"
)

var (
	// dbDNS e.g. postgres://username:password@url.com:5432/dbName
	dbDNS = ""
)

// Info acts as the DB model for storing the monitoring data
type Info struct {
	TimeOfExec          time.Time          `json:"time_of_exec"`
	ResponseTime        map[string]float64 `json:"response_time"`
	UnavailableServices []string           `json:"unavailable_services"`
}

// SetupDB will initiate DB connection and create the table if not exist
func SetupDB() error {
	dbDNS = fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.Data.DBUserName,
		config.Data.DBPassword,
		config.Data.DBHost,
		config.Data.DBPort,
		config.Data.DBName,
	)

	db, err := sql.Open("postgres", dbDNS)
	if err != nil {
		return fmt.Errorf("failed to open a DB connection: %v", err)
	}
	defer db.Close()

	query := `CREATE TABLE IF NOT EXISTS MONITOR (
		TIME_OF_EXEC TIMESTAMPTZ PRIMARY KEY,
		RESPONSE_TIME JSONB,
		UNAVAILABLE_SERVICES TEXT []
	)`

	_, err = db.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to execute query: %v", err)
	}

	return nil
}

// Save will save the data regarding the service connections to the DB
func (i *Info) Save() error {
	db, err := sql.Open("postgres", dbDNS)
	if err != nil {
		return fmt.Errorf("failed to open a DB connection: %v", err)
	}
	defer db.Close()
	jsonData, err := json.Marshal(&i.ResponseTime)
	if err != nil {
		return fmt.Errorf("error while trying to marshal data into bytes: %v", err)
	}
	_, err = db.Exec("INSERT INTO MONITOR (TIME_OF_EXEC,RESPONSE_TIME,UNAVAILABLE_SERVICES) VALUES($1,$2,$3)", i.TimeOfExec, jsonData, pq.Array(i.UnavailableServices))
	if err != nil {
		return err
	}
	return nil
}
