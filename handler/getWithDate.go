package handler

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"net/http"
)

// GetWithDate is for handling the search requests
func (c *Connector) GetWithDate(ctx iris.Context) {
	fromDate := ctx.URLParam("fromDate")
	fromTime := ctx.URLParam("fromTime")
	from := fmt.Sprintf("%s %s", fromDate, fromTime)
	if !validateDate(fromDate, fromTime) {
		ctx.StatusCode(http.StatusBadRequest)
		ctx.WriteString("incorrect format for from date or time")
		return
	}
	toDate := ctx.URLParam("toDate")
	toTime := ctx.URLParam("toTime")
	to := fmt.Sprintf("%s %s", toDate, toTime)
	if !validateDate(toDate, toTime) {
		ctx.StatusCode(http.StatusBadRequest)
		ctx.WriteString("incorrect format for to date or time")
		return
	}

	status := ctx.URLParam("status")
	if status != "" && (status != "available" && status != "unavailable") {
		ctx.StatusCode(http.StatusBadRequest)
		ctx.WriteString("status can only be available or unavailable")
		return
	}
	data, err := c.getByDate(from, to, status)
	if err != nil {
		ctx.StatusCode(http.StatusNotFound)
		ctx.WriteString(err.Error())
		return
	}

	ctx.StatusCode(http.StatusOK)
	ctx.JSON(data)
}
