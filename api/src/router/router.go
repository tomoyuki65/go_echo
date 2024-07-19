package router

import (
    "github.com/labstack/echo/v4"
    "api/internal/handlers/index"
    "api/internal/handlers/user"
    "api/internal/handlers/post"
    "api/internal/middleware/auth"
)

func SetupRouter(e *echo.Echo) {
    indexHandler := index.NewIndexHandler(e)
    e.GET("/", indexHandler.GetIndex)

    v1 := e.Group("/api/v1")
    v1.POST("/hello", indexHandler.PostIndex)

    userHandler := user.NewUserHandler(e)
    v1.POST("/user", userHandler.CreateUser)
    v1.GET("/user/:uid", userHandler.GetUser, auth.FirebaseAuth)
    v1.GET("/users", userHandler.GetUsers)
    v1.PUT("/user/:uid", userHandler.UpdateUser)
    v1.DELETE("/user/:uid", userHandler.DeleteUser)

    postHandler := post.NewPostHandler(e)
    v1.POST("/post", postHandler.CreatePost)
    v1.GET("/post/:id", postHandler.GetPost)
    v1.DELETE("/post/:id", postHandler.DeletePost)
}