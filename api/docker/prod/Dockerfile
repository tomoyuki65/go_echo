# ビルドステージ
FROM golang:1.22.3-alpine AS builder

WORKDIR /app

# 依存パッケージをキャッシュするため、go.mod と go.sum を先にコピー
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o main .

# 実行ステージ
FROM alpine:latest

WORKDIR /go

# ビルドステージで作成したバイナリをコピー
COPY --from=builder /app/main .

# ポートを設定
EXPOSE 8080

# APIサーバー起動コマンド
CMD ["./main"]