package service

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

type mockInfo struct {
	URL string
}

func (i *mockInfo) contactService() (*http.Response, error) {
	if i.URL == "invalid.com" {
		return &http.Response{
			StatusCode: http.StatusBadGateway,
		}, nil
	}
	if i.URL == "connection-refused.com" {
		return nil, fmt.Errorf("error: connection forcefully closed from the host side")
	}
	return &http.Response{
		StatusCode: http.StatusOK,
	}, nil
}

func TestMonitor(t *testing.T) {
	monitorChan := make(chan ExecStatus)
	type args struct {
		info connector
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Positive test case",
			args: args{
				info: &mockInfo{
					URL: "valid.com",
				},
			},
		},
		{
			name: "unavailable service",
			args: args{
				info: &mockInfo{
					URL: "invalid.com",
				},
			},
		},
		{
			name: "connection refused from host",
			args: args{
				info: &mockInfo{
					URL: "connection-refused.com",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			go Monitor(tt.args.info, monitorChan)
			x := mockWatch(monitorChan)
			infoMap := make(map[string]string)
			infoBytes, _ := json.Marshal(tt.args.info)
			json.Unmarshal(infoBytes, &infoMap)
			if infoMap["URL"] == "valid.com" {
				assert.Equal(t, true, x.Availability, fmt.Sprintf("Expected the availability of %v to be true", infoMap["URL"]))
			} else {
				assert.Equal(t, false, x.Availability, fmt.Sprintf("Expected the availability of %v to be false", infoMap["URL"]))
			}
		})
	}
}

func mockWatch(monitorChan chan ExecStatus) *ExecStatus {
	for {
		select {
		case x := <-monitorChan:
			return &x
		}
	}
}
