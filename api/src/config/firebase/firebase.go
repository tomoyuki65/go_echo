package firebase

import (
    "context"
    "fmt"

    "api/config"

    "firebase.google.com/go/v4/auth"
    "google.golang.org/api/option"
    firebase "firebase.google.com/go/v4"
)

func SetupFirebase() (*auth.Client, error) {
    // コンフィグ設定の取得
    cfg, err := config.SetupConfig("")
    if err != nil {
        fmt.Println("コンフィグ設定取得エラー", err)
        return nil, err
    }

    // Firebase接続設定
    opt := option.WithCredentialsJSON([]byte(fmt.Sprintf(
               `{
                   "type": "%s",
                   "project_id": "%s",
                   "private_key_id": "%s",
                   "private_key": %q,
                   "client_email": "%s",
                   "client_id": "%s",
                   "auth_uri": "%s",
                   "token_uri": "%s",
                   "auth_provider_x509_cert_url": "%s",
                   "client_x509_cert_url": "%s",
                   "universe_domain": "%s"
               }`,
               cfg.Fb.Type,
               cfg.Fb.ProjectId,
               cfg.Fb.PrivateKeyId,
               cfg.Fb.PrivateKey,
               cfg.Fb.ClientEmail,
               cfg.Fb.ClientId,
               cfg.Fb.AuthUri,
               cfg.Fb.TokenUri,
               cfg.Fb.AuthProviderX509CertUrl,
               cfg.Fb.ClientX509CertUrl,
               cfg.Fb.UniverseDomain,
           )))

    // FirebaseのAppを初期化
    app, err := firebase.NewApp(context.Background(), nil, opt)
    if err != nil {
        return nil, fmt.Errorf("error initializing app: %v", err)
    }

    // FirebaseのAuthクライアントを初期化
    client, err := app.Auth(context.Background())
    if err != nil {
        return nil, fmt.Errorf("error getting Auth client: %v", err)
    }

    return client, nil
}