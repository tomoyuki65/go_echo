##############################
# 本番環境用
##############################
####################
# ビルドステージ
####################
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

####################
# 実行ステージ
####################
FROM alpine:latest

WORKDIR /go/src

# インストール可能なパッケージ一覧の更新
RUN apk update && \
    apk upgrade && \
    # パッケージのインストール（--no-cacheでキャッシュ削除）
    apk add --no-cache \
            tzdata \
            curl

# ビルドステージで作成したバイナリをコピー
COPY --from=builder ./go/src/main .

# prod用の設定ファイルをコピー
COPY ./src/config/config.prod.yml ./config/

## 今回はコンテナデプロイ時にマイグレーションも実行させたいため、
## マイグレーション用ファイルのコピーとatlasコマンドのインストールも行う。
## （実務等では踏み台サーバーを用意してDBにアクセスさせるのでその場合は不要）
COPY ./src/ent ./ent
RUN curl -sSf https://atlasgo.sh | sh

# コンテナ起動時に実行させるシェルファイルをコピー
# シェルファイルからatlasコマンドでマイグレーションをできるようにする。
COPY ./docker/prod/entrypoint.sh .
RUN chmod +x entrypoint.sh

# ポートを設定
EXPOSE 8080

# シェルファイルを実行
ENTRYPOINT ["ash", "./entrypoint.sh"]

# APIサーバー起動コマンド
CMD ["./main"]