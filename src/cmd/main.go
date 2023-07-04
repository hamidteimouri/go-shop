package main

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"mytheresa/internal/data/memory"
	"mytheresa/internal/domain"
	"mytheresa/internal/presentation/httpserver"
	_ "mytheresa/pkg"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			logrus.WithFields(logrus.Fields{
				"err": r,
			}).Error("recovered")
		}
	}()

	// datasource
	ds := memory.NewDataSource()
	// domain
	productService := domain.NewProductService(ds)

	e := echo.New()
	e.HidePort = false
	e.HideBanner = true

	// registering routes and serving HTTP server
	httpserver.RegisterProductHandler(e, productService)
	go func() {
		err := e.Start(":8000")
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"err": err,
			}).Panic("failed to serve http server")
		}
	}()

	// wait for `Ctrl+c` or docker stop/restart signal to exit
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGKILL, syscall.SIGTERM)
	<-ch
}
