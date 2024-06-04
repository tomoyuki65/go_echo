package index

import (
    "testing"
    "api/internal/services"
    "net/http/httptest"
    "github.com/labstack/echo/v4"
	"api/pkg/validator"
    "github.com/stretchr/testify/assert"
    "net/http"
    "strings"
    "bytes"
)

// テスト用のEcho設定
func echoSetup() *echo.Echo {
	e := echo.New()
	e.Validator = validator.NewCustomValidator()

	return e
}

// メッセージ「{"message":"Hello, World !!"}」を出力して正常終了すること
func TestGetIndex(t *testing.T) {
	// ハンドラーのインスタンス作成
	handler := &indexHandler{
		helloService: services.NewHelloService(),
	}

	e := echoSetup()

	// テスト用サーバーの作成
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// ハンドラーの実行
	err := handler.GetIndex(c)
	if err != nil {
		c.Error(err) 
	}

    // 実行結果の検証
    assert.Equal(t, http.StatusOK, rec.Code)
    assert.Equal(t, `{"message":"Hello, World !!"}`, strings.TrimSpace(rec.Body.String()))	
}

// パラメータに指定したtextを出力して正常終了すること
func TestPostIndex(t *testing.T) {
	// ハンドラーのインスタンス作成
	handler := &indexHandler{
		helloService: services.NewHelloService(),
	}

	e := echoSetup()

	// テスト用サーバーの作成
	body := []byte(`{"text": "テスト実行！"}`)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/hello", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// ハンドラーの実行
	err := handler.PostIndex(c)
	if err != nil {
		c.Error(err) 
	}

    // 実行結果の検証
    assert.Equal(t, http.StatusOK, rec.Code)
    assert.Equal(t, `{"message":"テスト実行！"}`, strings.TrimSpace(rec.Body.String()))	
}

// パラメータにtextを指定しない場合に400エラーになること
func TestPostIndexNotParamText(t *testing.T) {
	// ハンドラーのインスタンス作成
	handler := &indexHandler{
		helloService: services.NewHelloService(),
	}

	e := echoSetup()

	// テスト用サーバーの作成
	req := httptest.NewRequest(http.MethodPost, "/api/v1/hello", nil)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// ハンドラーの実行
	err := handler.PostIndex(c)
	if err != nil {
		c.Error(err) 
	}

	// 実行結果の検証
	assert.Error(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}