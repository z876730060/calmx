package main

import (
	_ "calmx/docs"

	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

//	@title		gin Web App
//	@BasePath	/
func main() {

	e := echo.New()
	InitMiddleware(e)
	e.Use(echoprometheus.NewMiddleware("myapp"))   // adds middleware to gather metrics
	e.GET("/metrics", echoprometheus.NewHandler()) // adds route to serve gathered metrics

	e.GET("/hello", Index)
	e.POST("/login", Login)

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.Start(":1323")
}
