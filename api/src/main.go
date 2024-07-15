package main

import (
    "github.com/labstack/echo/v4"
    "api/router"
    "api/pkg/validator"
    echoSwagger "github.com/swaggo/echo-swagger"
    _ "api/docs"
)

// @title go_echo API
// @version 1.0
// @description Go言語（Golang）のフレームワーク「Echo」によるバックエンドAPIのサンプル
// @host localhost
// @BasePath /api/v1
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {
    e := echo.New()

    // API仕様書の設定
    e.GET("/swagger/*", echoSwagger.WrapHandler)

    // バリデーション設定
    e.Validator = validator.NewCustomValidator()

    // ルーティング設定
    router.SetupRouter(e)

    e.Logger.Fatal(e.Start(":8080"))
}