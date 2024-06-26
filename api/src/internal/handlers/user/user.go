package user

import (
    "context"
    "net/http"
    "time"

    "api/database"
    "api/internal/services/user"

    "github.com/labstack/echo/v4"
)

type CreateUserRequestBody struct {
    Uid string `json:"uid" form:"uid" validate:"required" example:"Ax1abc9uiowd9lKE"`
    LastName string `json:"last_name" form:"last_name" validate:"required" example:"田中"`
    FirstName string `json:"first_name" form:"first_name" validate:"required" example:"太郎"`
    Email string `json:"email" form:"email" validate:"required,email" example:"taro@example.com"`
}

type UpdateUserRequestBody struct {
    LastName string `json:"last_name" form:"last_name" example:"田中"`
    FirstName string `json:"first_name" form:"first_name" example:"太郎"`
    Email string `json:"email" form:"email" validate:"omitempty,email" example:"taro@example.com"`
}

type userHandler struct {
    userService user.UserService
}

func NewUserHandler(e *echo.Echo) *userHandler {
    return &userHandler{
        userService: user.NewUserService(),
    }
}

type MessageResponse struct {
    Message string `json:"message"`
}

type UserResponse struct {
    ID int64 `json:"id" example:"1"`
    UID string `json:"uid" example:"Xa12kK9ohsIhldD4"`
    LastName string `json:"last_name" example:"田中"`
    FirstName string `json:"first_name" example:"太郎"`
    Email string `json:"email" example:"taro@example.com"`
    CreatedAt time.Time `json:"created_at" example:"2024-06-23T23:18:49+09:00"`
    UpdatedAt time.Time `json:"updated_at" example:"2024-06-23T23:18:49+09:00"`
}

// @Description ユーザー作成
// @Tags user
// @Param CreateUser body CreateUserRequestBody true "作成するユーザー情報"
// @Success 201 {object} UserResponse ユーザー情報
// @Failure 400
// @Failure 500
// @Router /user [post]
func (h *userHandler) CreateUser(c echo.Context) error {
    var r CreateUserRequestBody
    if err := c.Bind(&r); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
    }

    // バリデーションを実行
    if err := c.Validate(&r); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }

    // DB設定
    dbCtx := context.Background()
    dbClient, err := database.SetupDatabase("")
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Failed Database Connection")
    }

    // ユーザー作成
    res, err2 := h.userService.CreateUser(
                     c.Request().Context(),
                     dbCtx,
                     dbClient,
                     r.Uid,
                     r.LastName,
                     r.FirstName,
                     r.Email,
                 )
    if err2 != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Failed Create User")
    }

    return c.JSON(http.StatusCreated, res)
}

// @Description 有効な対象ユーザー取得
// @Tags user
// @Param uid path string true "uid"
// @Success 200 {object} UserResponse ユーザー情報
// @Failure 404
// @Failure 405
// @Failure 500
// @Router /user/:uid [get]
func (h *userHandler) GetUser(c echo.Context) error {
    // パスパラメータ取得
    uid := c.Param("uid")

    // DB設定
    dbCtx := context.Background()
    dbClient, err := database.SetupDatabase("")
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Failed Database Connection")
    }

    // 対象ユーザー取得
    res, err2 := h.userService.GetUser(
                     c.Request().Context(),
                     dbCtx,
                     dbClient,
                     uid,
                 )
    if err2 != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, err2.Error())
    }

    return c.JSON(http.StatusOK, res)
}

// @Description 有効な全てのユーザー取得
// @Tags user
// @Success 200 {object} []UserResponse ユーザー情報
// @Failure 500
// @Router /users [get]
func (h *userHandler) GetUsers(c echo.Context) error {
    // DB設定
    dbCtx := context.Background()
    dbClient, err := database.SetupDatabase("")
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Failed Database Connection")
    }

    // 全てのユーザー取得
    res, err2 := h.userService.GetUsers(
                     c.Request().Context(),
                     dbCtx,
                     dbClient,
                 )
    if err2 != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, err2.Error())
    }

    return c.JSON(http.StatusOK, res)
}

// @Description 対象ユーザー更新
// @Tags user
// @Param UpdateUser body UpdateUserRequestBody true "更新するユーザー情報"
// @Success 200 {object} UserResponse ユーザー情報
// @Failure 404
// @Failure 405
// @Failure 500
// @Router /user/:uid [put]
func (h *userHandler) UpdateUser(c echo.Context) error {
    // パスパラメータ取得
    uid := c.Param("uid")

    var r UpdateUserRequestBody
    if err := c.Bind(&r); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
    }

    // バリデーションを実行
    if err := c.Validate(&r); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }

    // DB設定
    dbCtx := context.Background()
    dbClient, err := database.SetupDatabase("")
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Failed Database Connection")
    }

    // ユーザー更新
    res, err2 := h.userService.UpdateUser(
                     c.Request().Context(),
                     dbCtx,
                     dbClient,
                     uid,
                     r.LastName,
                     r.FirstName,
                     r.Email,
                 )
    if err2 != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, err2.Error())
    }

    return c.JSON(http.StatusOK, res)
}

// @Description 対象ユーザー削除
// @Tags user
// @Param uid path string true "uid"
// @Success 200 {object} MessageResponse メッセージ
// @Failure 404
// @Failure 405
// @Failure 500
// @Router /user/:uid [delete]
func (h *userHandler) DeleteUser(c echo.Context) error {
    // パスパラメータ取得
    uid := c.Param("uid")

    // DB設定
    dbCtx := context.Background()
    dbClient, err := database.SetupDatabase("")
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Failed Database Connection")
    }

    // ユーザー削除
    err2 := h.userService.DeleteUser(
                     c.Request().Context(),
                     dbCtx,
                     dbClient,
                     uid,
                 )
    if err2 != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, err2.Error())
    }

    res := MessageResponse{ Message: "OK" }

    return c.JSON(http.StatusOK, res)
}