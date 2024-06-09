FROM golang:1.22.3-alpine

WORKDIR /go/src

# インストール可能なパッケージ一覧の更新
RUN apk update && \
    apk upgrade && \
    # パッケージのインストール（--no-cacheでキャッシュ削除）
    apk add --no-cache \
            git

COPY ./src .

# モジュールの依存関係をダウンロード
RUN if [ -f ./go.mod ]; then \
      go mod download; \
    fi

# air をインストール
RUN go install github.com/air-verse/air@latest

EXPOSE 8080