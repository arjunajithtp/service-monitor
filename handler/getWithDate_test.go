package handler

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/httptest"
	"net/http"
	"testing"
)

func TestConnector_GetWithDate(t *testing.T) {
	c := mockGetConnector()
	expectedBody1 := `{"dd-mm-yyyy HH:MM:SS":["1.com","2.com"]}`
	expectedBody2 := `{"dd-mm-yyyy HH:MM:SS":["3.com","4.com"]}`
	app := iris.New()
	app.Get("/get-status", c.GetWithDate)
	e := httptest.New(t, app)

	e.GET(
		"/get-status",
	).WithQueryString(
		"fromDate=2020-11-15&fromTime=20:00:00&toDate=2020-11-16&toTime=20:02:00&status=available",
	).Expect().Status(http.StatusOK).Body().Equal(expectedBody1)

	e.GET(
		"/get-status",
	).WithQueryString(
		"fromDate=2020-11-15&fromTime=20:00:00&toDate=2020-11-16&toTime=20:02:00&timeTaken=less",
	).Expect().Status(http.StatusOK).Body().Equal(expectedBody2)

	e.GET(
		"/get-status",
	).WithQueryString(
		"fromDate=2020-&fromTime=20:00:00&toDate=2020-11-16&toTime=20:02:00&timeTaken=less",
	).Expect().Status(http.StatusBadRequest).Body().Equal("incorrect format for from date or time")

	e.GET(
		"/get-status",
	).WithQueryString(
		"fromDate=2020-11-15&fromTime=20:00:00&toDate=2020-11-16&toTime=200200&timeTaken=less",
	).Expect().Status(http.StatusBadRequest).Body().Equal("incorrect format for to date or time")

	e.GET(
		"/get-status",
	).WithQueryString(
		"fromDate=2020-11-15&fromTime=20:00:00&toDate=2020-11-16&toTime=20:02:00&status=abc",
	).Expect().Status(http.StatusBadRequest).Body().Equal("status can only be available or unavailable")

	e.GET(
		"/get-status",
	).WithQueryString(
		"fromDate=2020-11-15&fromTime=20:00:00&toDate=2020-11-16&toTime=20:02:00&timeTaken=abc",
	).Expect().Status(http.StatusBadRequest).Body().Equal("status can only be greater or less")
}

func mockGetByDate(from, to, status string) (map[string][]string, error) {
	return map[string][]string{
		"dd-mm-yyyy HH:MM:SS": []string{
			"1.com",
			"2.com",
		},
	}, nil
}

func mockGetByTimeTaken(from, to, timeTaken string) (map[string][]string, error) {
	return map[string][]string{
		"dd-mm-yyyy HH:MM:SS": []string{
			"3.com",
			"4.com",
		},
	}, nil
}

func mockGetConnector() *Connector {
	return &Connector{
		getByDate:      mockGetByDate,
		getByTimeTaken: mockGetByTimeTaken,
	}
}
