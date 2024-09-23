## ビルドステージ
FROM golang:1.23-alpine AS builder

WORKDIR /go/src

# インストール可能なパッケージ一覧の更新
RUN apk update && \
    apk upgrade && \
    # パッケージのインストール（--no-cacheでキャッシュ削除）
    apk add --no-cache \
            git \
            curl

COPY ./src .

# 依存関係を反映
RUN go install

# ビルド
RUN go build -o main .

## 実行ステージ
FROM alpine:latest

WORKDIR /go/src

# インストール可能なパッケージ一覧の更新
RUN apk update && \
    apk upgrade && \
    # パッケージのインストール（--no-cacheでキャッシュ削除）
    apk add --no-cache \
            tzdata \
            curl

# マイグレーション用のatlasをインストール
RUN curl -sSf https://atlasgo.sh | sh

# ビルドステージで作成したバイナリをコピー
COPY --from=builder ./go/src/main .

# マイグレーション用のファイルを全てコピー
COPY ./src/ent ./ent

# prod用設定ファイルをコピー
COPY ./src/config/config.prod.yml ./config/

# コンテナ起動時に実行するスクリプトファイルをコピー
COPY ./docker/prod/entrypoint.sh .
RUN chmod +x entrypoint.sh

# ポートを設定
EXPOSE 8080

# スクリプトを実行
ENTRYPOINT ["ash", "./entrypoint.sh"]

# APIサーバー起動コマンド
CMD ["./main"]