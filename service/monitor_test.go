package service

import (
	"net/http"
	"testing"
)

type mockInfo struct{}

func (i *mockInfo) contactService() (*http.Response, error) {
	return nil, nil
}

func (i *mockInfo) saveToDB(data string) error {
	return nil
}

func TestMonitor(t *testing.T) {
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
			Monitor(tt.args.info)
		})
	}
}
