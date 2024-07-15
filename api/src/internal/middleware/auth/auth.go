package auth

import (
    "context"
    "strings"
    "net/http"

    "api/config/firebase"

    "github.com/labstack/echo/v4"
)

func FirebaseAuth(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        // AuthorizationヘッダーからidTokenを取得
        idToken := ""
        authHeader := c.Request().Header.Get("Authorization")
        if authHeader != "" {
            idToken = strings.Replace(authHeader, "Bearer ", "", 1)
        } else {
            return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Unauthorized"})
        }

        // Firebaseのclient取得
        client, err := firebase.SetupFirebase()
        if err != nil {
            return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
        }

        // idTokenを検証
        token, err := client.VerifyIDToken(context.Background(), idToken)
        if err != nil {
            return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Unauthorized"})
        }

        // idTokenからuidを取得
        uid := token.UID

        // uidからユーザー情報取得
        firebaseUser, err := client.GetUser(context.Background(), uid)
        if err != nil {
            return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
        }

        // ユーザー情報からemailを取得
        email := firebaseUser.Email

        // uidとemailをContextに保存
        c.Set("uid", uid)
        c.Set("email", email)

        return next(c)
    }
}