FROM golang:1.23-alpine

WORKDIR /go/src

# インストール可能なパッケージ一覧の更新
RUN apk update && \
    apk upgrade && \
    # パッケージのインストール（--no-cacheでキャッシュ削除）
    apk add --no-cache \
            git \
            curl

COPY ./src .

# モジュールの依存関係をインストール
RUN if [ -f ./go.mod ]; then \
      go install; \
    fi

# air をインストール
RUN go install github.com/air-verse/air@latest

# API仕様書作成用のswagをインストール
RUN go install github.com/swaggo/swag/cmd/swag@latest

# マイグレーション用のatlasをインストール
RUN curl -sSf https://atlasgo.sh | sh

EXPOSE 8080