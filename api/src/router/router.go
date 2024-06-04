package router

import (
    "github.com/labstack/echo/v4"
    "api/internal/handlers/index"
)

func SetupRouter(e *echo.Echo) {
    indexHandler := index.NewIndexHandler(e)
    e.GET("/", indexHandler.GetIndex)

    v1 := e.Group("/api/v1")
    v1.POST("/hello", indexHandler.PostIndex)
}