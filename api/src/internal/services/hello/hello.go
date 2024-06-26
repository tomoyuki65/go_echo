package hello

import (
    "context"
)

// インターフェースの定義
type HelloService interface {
    GetHello(ctx context.Context) (string, error)
    PostHello(ctx context.Context, text string) (string, error)
}

// 空の構造体を定義
type helloService struct{}

// helloServiceのポインタを返す関数
func NewHelloService() HelloService {
    return &helloService{}
}

// インターフェースの各サービスを実装
func (s *helloService) GetHello(ctx context.Context) (string, error) {

    res := "Hello, World !!"

    return res, nil
}

func (s *helloService) PostHello(ctx context.Context, text string) (string, error) {

    res := "Hello, World !!"

    if (text != "") {
        res = text
    }

    return res, nil
}