package service

import (
	"encoding/json"
	"fmt"
	"github.com/arjunajithtp/service-monitor/config"
	"github.com/arjunajithtp/service-monitor/model"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var counter int

func (i *mockInfo) saveToDB(data model.Info) error {
	if data.ResponseTime[i.URL] == 2.2 {
		counter++
	}
	return nil
}

func TestWatch(t *testing.T) {
	config.Data.Services = append(config.Data.Services, "save.com")
	monitorChan := make(chan ExecStatus)
	type args struct {
		c connector
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "saveDataToDB gets correct data",
			args: args{
				c: &mockInfo{
					URL: "save.com",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			infoMap := make(map[string]string)
			infoBytes, _ := json.Marshal(tt.args.c)
			json.Unmarshal(infoBytes, &infoMap)
			go mockMonitor(infoMap["URL"], monitorChan)
			Watch(tt.args.c, monitorChan)
			time.Sleep(2 * time.Second)
			assert.Equal(t, 1, counter, fmt.Sprintf("Expected the counter to be %v but got %v", 1, counter))
		})
	}
}

func mockMonitor(service string, monitorChan chan ExecStatus) {
	monitorChan <- ExecStatus{Service: service, ElapsedTime: 2.2, Availability: true}
}
