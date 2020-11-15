package main

import (
	"github.com/arjunajithtp/service-monitor/config"
	"github.com/arjunajithtp/service-monitor/service"
	"github.com/kataras/iris/v12"
	"log"
	"time"
)

func main() {
	if err := config.SetConfiguration(); err != nil {
		log.Fatalf("error while trying to read the config file: %v", err)
	}

	ticker := time.NewTicker(time.Duration(config.Data.MonitoringIntervalInSec) * time.Second)
	go func() {
		for ; true; <-ticker.C {
			currentTime := time.Now()
			for _, url := range config.Data.Services {
				serviceInfo := &service.Info{
					URL:  url,
					Time: currentTime,
				}
				go service.Monitor(serviceInfo)
			}
		}
	}()

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
