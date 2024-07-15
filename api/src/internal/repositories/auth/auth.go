package auth

import (
    "context"
    "fmt"

    "api/config/firebase"

    "firebase.google.com/go/v4/auth"
)

func UpdateUser(
    uid string,
    email string,
    password string,
) (*auth.UserRecord, error) {
    // Firebaseのclient取得
    client, err := firebase.SetupFirebase()
    if err != nil {
        return nil, fmt.Errorf("failed get firebase client: %v", err)
    }

    // Firebaseユーザー更新
    params := &auth.UserToUpdate{}
    if email != "" {
        params = params.Email(email)
    }
    if password != "" {
        params = params.Password(password)
    }
    user, err := client.UpdateUser(context.Background(), uid, params)
    if err != nil {
        return nil, fmt.Errorf("failed update firebase user: %v", err)
    }

    return user, nil
}

func DeleteUser(uid string) error {
    // Firebaseのclient取得
    client, err := firebase.SetupFirebase()
    if err != nil {
        return fmt.Errorf("failed get firebase client: %v", err)
    }

    // Firebaseユーザー削除
    err = client.DeleteUser(context.Background(), uid)
    if err != nil {
        return fmt.Errorf("failed delete firebase user: %v", err)
    }

    return nil
}