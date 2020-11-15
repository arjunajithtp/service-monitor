package service

import (
	"github.com/arjunajithtp/service-monitor/model"
	"net/http"
	"testing"
	"time"
)

type mockInfo struct {
	Time time.Time
	URL  string
}

func (i *mockInfo) contactService() (*http.Response, error) {
	return nil, nil
}

func (i *mockInfo) saveToDB(data model.Info) error {
	return nil
}

func TestMonitor(t *testing.T) {
	monitorChan := make(chan ExecStatus)
	info := mockInfo{}
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
				info: &info,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Monitor(tt.args.info, monitorChan)
		})
	}
}
