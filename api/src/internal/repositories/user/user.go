package user

import (
    "context"
    "fmt"
    "time"

    "api/ent"
    entUser "api/ent/user"
)

func CreateUser(
    dbCtx context.Context,
    dbClient *ent.Client,
    uid string,
    lastName string,
    firstName string,
    email string,
) (*ent.User, error) {

    user, err := dbClient.User.
                          Create().
                          SetUID(uid).
                          SetLastName(lastName).
                          SetFirstName(firstName).
                          SetEmail(email).
                          Save(dbCtx)
    if err != nil {
        return nil, fmt.Errorf("failed creating user: %v", err)
    }

    return user, nil
}

func GetUserFromEmail(
    dbCtx context.Context,
    dbClient *ent.Client,
    email string,
) (*ent.User, error) {

    user, err := dbClient.User.
                          Query().
                          Where(entUser.EmailEQ(email)).
                          Where(entUser.DeletedAtIsNil()).
                          First(dbCtx)
    if err != nil && !ent.IsNotFound(err) {
        return nil, fmt.Errorf("failed get user: %v", err)
    }

    return user, nil
}

func GetUserFromUid(
    dbCtx context.Context,
    dbClient *ent.Client,
    uid string,
) (*ent.User, error) {

    user, err := dbClient.User.
                          Query().
                          Where(entUser.UIDEQ(uid)).
                          Where(entUser.DeletedAtIsNil()).
                          First(dbCtx)
    if err != nil && !ent.IsNotFound(err) {
        return nil, fmt.Errorf("failed get user: %v", err)
    }

    return user, nil
}

func GetUsers(
    dbCtx context.Context,
    dbClient *ent.Client,
) ([]*ent.User, error) {

    users, err := dbClient.User.
                           Query().
                           Where(entUser.DeletedAtIsNil()).
                           All(dbCtx)
    if err != nil {
        return nil, fmt.Errorf("failed get users: %v", err)
    }

    return users, nil
}

func UpdateUser(
    dbCtx context.Context,
    dbClient *ent.Client,
    user *ent.User,
    lastName string,
    firstName string,
    email string,
) (*ent.User, error) {

    updateOneUser := dbClient.User.
                              UpdateOne(user)

    if lastName != "" {
        updateOneUser = updateOneUser.SetLastName(lastName)
    }

    if firstName != "" {
        updateOneUser = updateOneUser.SetFirstName(firstName)
    }

    if email != "" {
        updateOneUser = updateOneUser.SetEmail(email)
    }

    user, err := updateOneUser.Save(dbCtx)
    if err != nil {
        return nil, fmt.Errorf("failed update user: %v", err)
    }

    return user, nil
}

func DeleteUser(
    dbCtx context.Context,
    dbClient *ent.Client,
    user *ent.User,
) error {
    // 現在の日時を文字列で取得
    date := time.Now()
    dateString := date.Format("2006-01-02 15:04:05")

    // 更新用のemailの値を設定
    updateEmail := user.Email + dateString

    _, err := dbClient.User.
                       UpdateOne(user).
                       SetEmail(updateEmail).
                       SetDeletedAt(time.Now()).
                       Save(dbCtx)
    if err != nil {
        return fmt.Errorf("failed delete user: %v", err)
    }

    return nil
}