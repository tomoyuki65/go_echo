package config

import (
    "os"
    "github.com/spf13/viper"
    "github.com/fsnotify/fsnotify"
    "fmt"
)

type Config struct {
    Env string
    Db  Database `yml:db`
}

type Database struct {
    UserName  string
    Password  string
    Host      string `yml:host`
    Port      int    `yml:port`
    Database  string `yml:database`
    Charset   string `yml:charset`
    Collation string `yml:collation`
    ParseTime bool   `yml:parseTime`
    Loc       string
}

func SetupConfig(env string) (*Config, error) {

    var config *Config

    if env != "testing" {
        env = os.Getenv("ENV")
    } 

    user := os.Getenv("MYSQL_USER")
    password := os.Getenv("MYSQL_PASSWORD")
    tz := os.Getenv("TZ")

    // viperの初期設定
    viper.SetConfigName("config." + env)
    viper.SetConfigType("yml")
    viper.AddConfigPath("config/")

    // configファイルの読み込み
    err := viper.ReadInConfig()
    if err != nil {
      panic(err)
    }

    // 読み込んだデータを変数cfgに設定
    err = viper.Unmarshal(&config)
    if err != nil {
      panic(err)
    }

    // configファイル更新時に再読み込み
    viper.WatchConfig()
    viper.OnConfigChange(func(e fsnotify.Event) {
        fmt.Println("Config file changed:", e.Name)
    })

    // 環境変数の値を変数cfgに設定
    config.Env = env
    config.Db.UserName = user
    config.Db.Password = password
    config.Db.Loc = tz

    return config, nil
}
