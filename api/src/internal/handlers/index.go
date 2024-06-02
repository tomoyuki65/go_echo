package handlers

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
type helloHandler struct {
	helloService services.HelloService
}

// helloHandlerのポインタを返す関数
func NewHelloHandler(e *echo.Echo) *helloHandler {
	return &helloHandler{
		helloService: services.NewHelloService(),
	}
}

// ハンドラーを実装
func (h *helloHandler) GetIndex(c echo.Context) error {
	getHello, err := h.helloService.GetHello(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get hello")
	}

	res := map[string]interface{}{
		       "message": getHello,
	       }

	return c.JSON(http.StatusOK, res)
}

func (h *helloHandler) PostIndex(c echo.Context) error {
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

	res := map[string]interface{}{
		       "message": postHello,
	       }

	return c.JSON(http.StatusOK, res)
}