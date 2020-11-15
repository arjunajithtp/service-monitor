package service

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// Monitor will monitor the avialbility of the services provided in the config file
func Monitor(c connector, monitorChan chan ExecStatus) {
	//var execStatus ExecStatus
	infoMap := make(map[string]string)
	infoBytes, err := json.Marshal(c)
	if err != nil {
		log.Printf("error while trying to convert the connector interface data into bytes: %v", err)
		return
	}
	err = json.Unmarshal(infoBytes, &infoMap)
	if err != nil {
		log.Printf("error while trying to convert the connector interface data into map: %v", err)
		return
	}

	start := time.Now()
	resp, err := c.contactService()
	elapsed := time.Since(start).Seconds()

	if err != nil || resp == nil || resp.StatusCode == http.StatusBadGateway {
		log.Printf("error: got %v while trying to connect with the service with url %v: %v", resp, infoMap["URL"], err)
		monitorChan <- ExecStatus{Service: infoMap["URL"], ElapsedTime: elapsed, Availability: false}
		return
	}

	monitorChan <- ExecStatus{Service: infoMap["URL"], ElapsedTime: elapsed, Availability: true}

}
