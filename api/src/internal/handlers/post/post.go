package post

import (
    "context"
    "net/http"
    "time"
    "strconv"

    "api/database"
    "api/internal/services/post"

    "github.com/labstack/echo/v4"
)

type CreatePostRequestBody struct {
    UserId int64 `json:"user_id" form:"user_id" validate:"required" example:"1"`
    Text string `json:"text" form:"text" validate:"required" example:"こんんちは。"`
}

type postHandler struct {
    postService post.PostService
}

func NewPostHandler(e *echo.Echo) *postHandler {
    return &postHandler{
        postService: post.NewPostService(),
    }
}

type MessageResponse struct {
    Message string `json:"message"`
}

type PostResponse struct {
    ID int64 `json:"id" example:"1"`
    UserId int64 `json:"user_id" example:"1"`
    Text string `json:"text" example:"こんにちは。"`
    CreatedAt time.Time `json:"created_at" example:"2024-06-23T23:18:49+09:00"`
    UpdatedAt time.Time `json:"updated_at" example:"2024-06-23T23:18:49+09:00"`
    Edges struct {}
}

type User struct {
    ID int64 `json:"id" example:"1"`
    UID string `json:"uid" example:"Xa12kK9ohsIhldD4"`
    LastName string `json:"last_name" example:"田中"`
    FirstName string `json:"first_name" example:"太郎"`
    Email string `json:"email" example:"taro@example.com"`
    CreatedAt time.Time `json:"created_at" example:"2024-06-23T23:18:49+09:00"`
    UpdatedAt time.Time `json:"updated_at" example:"2024-06-23T23:18:49+09:00"`
    Edges struct {}
}

type PostResponseWithUser struct {
    ID int64 `json:"id" example:"1"`
    UserId int64 `json:"user_id" example:"1"`
    Text string `json:"text" example:"こんにちは。"`
    CreatedAt time.Time `json:"created_at" example:"2024-06-23T23:18:49+09:00"`
    UpdatedAt time.Time `json:"updated_at" example:"2024-06-23T23:18:49+09:00"`
    Edges struct {
        Users User `json:"users"`
    } `json:"edges"`
}

// @Description Post作成
// @Tags post
// @Param CreateUser body CreatePostRequestBody true "作成するPost情報"
// @Success 201 {object} PostResponse ポスト情報
// @Failure 400
// @Failure 500
// @Router /post [post]
func (h *postHandler) CreatePost(c echo.Context) error {
    var r CreatePostRequestBody
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

    // Post作成
    res, err := h.postService.CreatePost(
                     c.Request().Context(),
                     dbCtx,
                     dbClient,
                     r.UserId,
                     r.Text,
                 )
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Failed Create Post")
    }

    return c.JSON(http.StatusCreated, res)
}

// @Description 有効な対象Post取得
// @Tags post
// @Param id path int64 true "id"
// @Success 200 {object} PostResponseWithUser ポスト情報（User含む）
// @Failure 404
// @Failure 405
// @Failure 500
// @Router /post/{id} [get]
func (h *postHandler) GetPost(c echo.Context) error {
    // パスパラメータ取得
    idStr := c.Param("id")
    id, err := strconv.ParseInt(idStr, 10, 64)
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Failed Id Conversion")
    }

    // DB設定
    dbCtx := context.Background()
    dbClient, err := database.SetupDatabase("")
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Failed Database Connection")
    }

    // 対象Post取得
    res, err := h.postService.GetPost(
                    c.Request().Context(),
                    dbCtx,
                    dbClient,
                    id,
                 )
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
    }

    return c.JSON(http.StatusOK, res)
}

// @Description 対象Post削除
// @Tags post
// @Param id path int64 true "id"
// @Success 200 {object} MessageResponse メッセージ
// @Failure 404
// @Failure 405
// @Failure 500
// @Router /post/{id} [delete]
func (h *postHandler) DeletePost(c echo.Context) error {
    // パスパラメータ取得
    idStr := c.Param("id")
    id, err := strconv.ParseInt(idStr, 10, 64)
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Failed Id Conversion")
    }

    // DB設定
    dbCtx := context.Background()
    dbClient, err := database.SetupDatabase("")
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Failed Database Connection")
    }

    // Post削除
    err = h.postService.DeletePost(
              c.Request().Context(),
              dbCtx,
              dbClient,
              id,
          )
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
    }

    res := MessageResponse{ Message: "OK" }

    return c.JSON(http.StatusOK, res)
}