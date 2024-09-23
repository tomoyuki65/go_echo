#!/bin/bash

# 環境変数「EXECUTE_MIGRATION」の値で制御する
if [ "$EXECUTE_MIGRATION" == "true" ]; then
  echo "マイグレーション実行"
  atlas migrate apply \
  --dir "file://ent/migrate/migrations" \
  --url "mysql://$MYSQL_USER:$MYSQL_PASSWORD@$MYSQL_HOST:3306/$MYSQL_DATABASE"
fi

# DockerfileのCMDを実行
exec "$@"