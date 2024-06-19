package database

import (
    "fmt"
    "net/url"
    "context"
    "log"

    "api/config"
    "api/ent"

    "entgo.io/ent/dialect"
    _ "github.com/go-sql-driver/mysql"
)

func SetupDatabase(env string) (*ent.Client, error) {

    // コンフィグ設定の取得
    var cfg *config.Config
    var err error

    if env == "testing" {
        cfg, err = config.SetupConfig("testing")
    } else {
        cfg, err = config.SetupConfig("")
    }

    if err != nil {
        fmt.Println("コンフィグ設定取得エラー", err)
        return nil, err
    }

    // ロケーションのエンコード
    encodedLoc := url.PathEscape(cfg.Db.Loc)

    dsn := fmt.Sprintf(
        "%s:%s@tcp(%s:%d)/%s?charset=%s&collation=%s&parseTime=%t&loc=%s",
        cfg.Db.UserName,
        cfg.Db.Password,
        cfg.Db.Host,
        cfg.Db.Port,
        cfg.Db.Database,
        cfg.Db.Charset,
        cfg.Db.Collation,
        cfg.Db.ParseTime,
        encodedLoc,
    )

    // DB接続
    var client *ent.Client
    client, err = ent.Open(dialect.MySQL, dsn)

    if err != nil {
        fmt.Println("DB接続エラー", err)
        return nil, err
    }

    // DBスキーマ作成
	if err = client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("DBスキーマ作成に失敗しました。: %v", err)
	}

    return client, nil
}

func GetDsnForMigrate(env string) (string, error) {

    // コンフィグ設定の取得
    var cfg *config.Config
    var err error

    if env == "testing" {
        cfg, err = config.SetupConfig("testing")
    } else {
        cfg, err = config.SetupConfig("")
    }

    if err != nil {
        fmt.Println("コンフィグ設定取得エラー", err)
        return "", err
    }

    // DBの接続先設定
    dsn := fmt.Sprintf(
               "mysql://%s:%s@%s:%d/%s?charset=%s&collation=%s&parseTime=%t&loc=%s",
               cfg.Db.UserName,
               cfg.Db.Password,
               cfg.Db.Host,
               cfg.Db.Port,
               cfg.Db.Database,
               cfg.Db.Charset,
               cfg.Db.Collation,
               cfg.Db.ParseTime,
               cfg.Db.Loc,
           )

    return dsn, nil
}