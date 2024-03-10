package server

import (
	"efficientDevelopment/db"
	"efficientDevelopment/internal/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"time"
)

type Config struct {
	Database   db.Database
	ServerPort int `server:"SERVER_PORT" default:"80"`
}

func Init() {

	e := echo.New()

	e.Use(middleware.CORS())

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
	})

	e.GET("/get-data", service.GetData)

	s := &http.Server{
		Addr:         ":8099",
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	e.Logger.Fatal(e.StartServer(s))

}
