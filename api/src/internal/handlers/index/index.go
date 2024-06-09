package index

import (
    "net/http"
    "github.com/labstack/echo/v4"
    "api/internal/services"
)

// リクエストボディの構造体
type RequestBody struct {
    Text string `json:"text" form:"text" validate:"required"`
}

// ハンドラーの構造体
type indexHandler struct {
    helloService services.HelloService
}

// indexHandlerのポインタを返す関数
func NewIndexHandler(e *echo.Echo) *indexHandler {
    return &indexHandler{
        helloService: services.NewHelloService(),
    }
}

// PostIndexのレスポンスの構造体
type PostIndexResponse struct {
    Message string `json:"message"`
}

// ハンドラーを実装
func (h *indexHandler) GetIndex(c echo.Context) error {
    getHello, err := h.helloService.GetHello(c.Request().Context())
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get hello")
    }

    res := map[string]interface{}{
               "message": getHello,
           }

    return c.JSON(http.StatusOK, res)
}

func (h *indexHandler) PostIndex(c echo.Context) error {
    var r RequestBody
    if err := c.Bind(&r); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
    }

    // バリデーションを実行
    if err := c.Validate(&r); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }

    postHello, err := h.helloService.PostHello(c.Request().Context(), r.Text)
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Failed to post hello")
    }
    
    res := PostIndexResponse{
        Message: postHello,
    }

    return c.JSON(http.StatusOK, res)
}