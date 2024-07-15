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
    Fb  Firebase
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

type Firebase struct {
    Type                    string
    ProjectId               string
    PrivateKeyId            string
    PrivateKey              string
    ClientEmail             string
    ClientId                string
    AuthUri                 string
    TokenUri                string
    AuthProviderX509CertUrl string
    ClientX509CertUrl       string
    UniverseDomain          string
}

func SetupConfig(env string) (*Config, error) {

    var config *Config

    if env != "testing" {
        env = os.Getenv("ENV")
    } 

    user := os.Getenv("MYSQL_USER")
    password := os.Getenv("MYSQL_PASSWORD")
    tz := os.Getenv("TZ")

    // Firebase
    fbType :=  os.Getenv("FIREBASE_TYPE")
    fbProjectId := os.Getenv("FIREBASE_PROJECT_ID")
    fbPrivateKeyId := os.Getenv("FIREBASE_PRIVATE_KEY_ID")
    fbPrivateKey := os.Getenv("FIREBASE_PRIVATE_KEY")
    fbClientEmail := os.Getenv("FIREBASE_CLIENT_EMAIL")
    fbClientId := os.Getenv("FIREBASE_CLIENT_ID")
    fbAuthUri := os.Getenv("FIREBASE_AUTH_URI")
    fbTokenUri := os.Getenv("FIREBASE_TOKEN_URI")
    fbAuthProviderX509CertUrl := os.Getenv("FIREBASE_AUTH_PROVIDER_X509_CERT_URL")
    fbClientX509CertUrl := os.Getenv("FIREBASE_CLIENT_X509_CERT_URL")
    fbUniverseDomain := os.Getenv("FIREBASE_UNIVERSE_DOMAIN")

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

    // Firebase
    config.Fb.Type = fbType
    config.Fb.ProjectId = fbProjectId
    config.Fb.PrivateKeyId = fbPrivateKeyId
    config.Fb.PrivateKey = fbPrivateKey
    config.Fb.ClientEmail = fbClientEmail
    config.Fb.ClientId = fbClientId
    config.Fb.AuthUri = fbAuthUri
    config.Fb.TokenUri = fbTokenUri
    config.Fb.AuthProviderX509CertUrl = fbAuthProviderX509CertUrl
    config.Fb.ClientX509CertUrl = fbClientX509CertUrl
    config.Fb.UniverseDomain = fbUniverseDomain

    return config, nil
}
