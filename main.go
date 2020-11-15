package main

import (
	"github.com/arjunajithtp/service-monitor/config"
	"github.com/arjunajithtp/service-monitor/model"
	"github.com/arjunajithtp/service-monitor/service"
	"github.com/kataras/iris/v12"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func init() {
	if err := config.SetConfiguration(); err != nil {
		log.Fatalf("error while trying to read the config file: %v", err)
	}

	model.SetupDB()
}

func main() {
	go initiateMonitoring()

	app := newApp()
	app.Run(iris.Addr(":"+config.Data.Port), iris.WithoutServerError(iris.ErrServerClosed))
}

func newApp() *iris.Application {
	app := iris.New()
	app.Get("/publish", handler)

	return app
}

func handler(ctx iris.Context) {

}

func initiateMonitoring() {
	ticker := time.NewTicker(time.Duration(config.Data.MonitoringIntervalInMin) * time.Minute)
	monitorChan := make(chan service.ExecStatus)
	for ; true; <-ticker.C {
		go service.Watch(&service.Info{}, monitorChan)
		for _, url := range config.Data.Services {
			serviceInfo := &service.Info{
				URL: url,
			}
			go service.Monitor(serviceInfo, monitorChan)
		}
	}
}
