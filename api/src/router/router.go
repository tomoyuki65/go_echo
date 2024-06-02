package router

import (
	"github.com/labstack/echo/v4"
	"api/internal/handlers"
)

func SetupRouter(e *echo.Echo) {
	helloHandler := handlers.NewHelloHandler(e)
	e.GET("/", helloHandler.GetIndex)

	v1 := e.Group("/api/v1")
    v1.POST("/hello", helloHandler.PostIndex)
}