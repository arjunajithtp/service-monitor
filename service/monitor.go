package service

import (
	"net/http"
	"time"
)

type connector interface {
	contactService() (*http.Response, error)
	saveToDB(string) error
}

// Info holds the information required for contacting the service
type Info struct {
	Time time.Time
	URL  string
}

// Monitor will monitor the avialbility of the services provided in the config file
func Monitor(info connector) {

}

// contactService contacts the web-services with the available info
// and returns the response and error if any
func (i *Info) contactService() (*http.Response, error) {
	return nil, nil
}

// saveToDB saves the service response information to the DB
// and it returns error if any
func (i *Info) saveToDB(data string) error {
	return nil
}
