package service

import (
	"encoding/json"
	"log"
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
	infoMap := make(map[string]interface{})
	infoBytes, err := json.Marshal(info)
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
	resp, err := info.contactService()
	elapsed := time.Since(start).Seconds()

	if err != nil {
		log.Printf("error: got %v while trying to connect with the service with url %v: %v", resp, infoMap["URL"], err)
	}

	log.Printf("http.Get to %s took %v seconds \n", infoMap["URL"], elapsed)
}

// contactService contacts the web-services with the available info
// and returns the response and error if any
func (i *Info) contactService() (*http.Response, error) {

	resp, err := http.Get("http://"+i.URL)

	if err != nil {
		return resp, err
	}

	defer resp.Body.Close()
	return resp, nil
}

// saveToDB saves the service response information to the DB
// and it returns error if any
func (i *Info) saveToDB(data string) error {
	return nil
}
