package main

import (
    "github.com/labstack/echo/v4"
    "api/router"
    "api/pkg/validator"
)

func main() {
    e := echo.New()

    // バリデーション設定
    e.Validator = validator.NewCustomValidator()

    // ルーティング設定
    router.SetupRouter(e)

    e.Logger.Fatal(e.Start(":8080"))
}