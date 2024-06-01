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
RUN go mod download

# air をインストール
RUN go install github.com/cosmtrek/air@latest

EXPOSE 8080

# air でホットリロードしながら起動
CMD ["air", "-c", ".air.toml"]