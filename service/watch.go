package service

import (
	"github.com/arjunajithtp/service-monitor/config"
	"github.com/arjunajithtp/service-monitor/model"
	"log"
	"time"
)

// Watch will listern to the monitorChan for all the service status
// and respond back with the result of the connections
func Watch(c connector, monitorChan chan ExecStatus) {
	var info model.Info
	info.ResponseTime = make(map[string]float64)
	info.TimeOfExec = time.Now()
	for {
		select {
		case x := <-monitorChan:
			if x.Availability {
				info.ResponseTime[x.Service] = x.ElapsedTime
			} else {
				info.UnavailableServices = append(info.UnavailableServices, x.Service)
			}
		}

		if len(info.ResponseTime)+len(info.UnavailableServices) == len(config.Data.Services) {
			err := c.saveToDB(info)
			if err != nil {
				log.Printf("error while trying to save data to the DB: %v", err)
			}
			return
		}

	}
}
