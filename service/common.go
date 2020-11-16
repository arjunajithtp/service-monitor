package service

import (
	"fmt"
	"github.com/arjunajithtp/service-monitor/model"
	"net/http"
)

type connector interface {
	contactService() (*http.Response, error)
	saveToDB(model.Info) error
}

// Info holds the information required for contacting the service
type Info struct {
	URL string
}

// ExecStatus holds the information required for contacting the service
type ExecStatus struct {
	Service      string
	ElapsedTime  float64
	Availability bool
}

// contactService contacts the web-services with the available info
// and returns the response and error if any
func (i *Info) contactService() (*http.Response, error) {

	resp, err := http.Get("http://" + i.URL)

	if err != nil {
		return resp, err
	}

	defer resp.Body.Close()
	return resp, nil
}

// saveToDB saves the service response information to the DB
// and it returns error if any
func (i *Info) saveToDB(data model.Info) error {
	err := data.Save()
	if err != nil {
		return fmt.Errorf("error while trying to insert into DB: %v", err)
	}
	return nil
}
