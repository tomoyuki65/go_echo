package user

import (
    "context"
    "fmt"

    "api/ent"
    "api/internal/repositories/user"
    "api/internal/repositories/auth"
)

type UserService interface {
    CreateUser(
        echoCtx context.Context,
        dbCtx context.Context,
        dbClient *ent.Client,
        uid string,
        lastName string,
        firstName string,
        email string,
    ) (*ent.User, error)
    GetUser(
        echoCtx context.Context,
        dbCtx context.Context,
        dbClient *ent.Client,
        email string,
    ) (UserResponse, error)
    GetUsers(
        echoCtx context.Context,
        dbCtx context.Context,
        dbClient *ent.Client,
    ) ([]*ent.User, error)
    UpdateUser(
        echoCtx context.Context,
        dbCtx context.Context,
        dbClient *ent.Client,
        uid string,
        lastName string,
        firstName string,
        email string,
        password string,
    ) (UserResponse, error)
    DeleteUser(
        echoCtx context.Context,
        dbCtx context.Context,
        dbClient *ent.Client,
        uid string,
    ) error
}

type userService struct{}

func NewUserService() UserService {
    return &userService{}
}

// レスポンス定義
type User struct { *ent.User }
type EmptyResponse map[string]interface{}
type UserResponse interface {}
func (u User) UserResponse() {}
func (e EmptyResponse) UserResponse() {}

func (s *userService) CreateUser(
    echoCtx context.Context,
    dbCtx context.Context,
    dbClient *ent.Client,
    uid string,
    lastName string,
    firstName string,
    email string,
) (*ent.User, error) {

    createUser, err := user.CreateUser(
                            dbCtx,
                            dbClient,
                            uid,
                            lastName,
                            firstName,
                            email,
                       )
    if err != nil {
        return nil, err
    }

    return createUser, nil
}

func (s *userService) GetUser(
    echoCtx context.Context,
    dbCtx context.Context,
    dbClient *ent.Client,
    uid string,
) (UserResponse, error) {

    getUser, err := user.GetUserFromUid(
                        dbCtx,
                        dbClient,
                        uid,
                    )
    if err != nil {
        return nil, err
    }

    if getUser == nil {
        emptyResponse := map[string]interface{}{}
        return emptyResponse, nil
    }

    return getUser, nil
}

func (s *userService) GetUsers(
    echoCtx context.Context,
    dbCtx context.Context,
    dbClient *ent.Client,
) ([]*ent.User, error) {

    getUsers, err := user.GetUsers(
                         dbCtx,
                         dbClient,
                     )
    if err != nil {
        return nil, err
    }

    return getUsers, nil
}

func (s *userService) UpdateUser(
    echoCtx context.Context,
    dbCtx context.Context,
    dbClient *ent.Client,
    uid string,
    lastName string,
    firstName string,
    email string,
    password string,
) (UserResponse, error) {

    getUser, err := user.GetUserFromUid(
                        dbCtx,
                        dbClient,
                        uid,
                    )
    if err != nil {
        return nil, err
    }

    if getUser == nil {
        return nil, fmt.Errorf("no user")
    }

    // トランザクション設定
    txDbClient, err := dbClient.Tx(dbCtx)
    if err != nil {
        return nil, err
    }

    updateUser, err := user.UpdateUser(
                           dbCtx,
                           txDbClient,
                           getUser,
                           lastName,
                           firstName,
                           email,
                       )
    if err != nil {
        // DBロールバック
        txDbClient.Rollback()

        return nil, err
    }

    if email != "" || password != "" {
        // Firebaseユーザー更新
        _, err := auth.UpdateUser(uid, email, password)
        if err != nil {
            // DBロールバック
            txDbClient.Rollback()

            return nil, err
        }
    }

    // DBコミット
    txDbClient.Commit()

    return updateUser, nil
}

func (s *userService) DeleteUser(
    echoCtx context.Context,
    dbCtx context.Context,
    dbClient *ent.Client,
    uid string,
) error {

    getUser, err := user.GetUserFromUid(
                        dbCtx,
                        dbClient,
                        uid,
                    )
    if err != nil {
        return err
    }

    if getUser == nil {
        return fmt.Errorf("no user")
    }

    // トランザクション設定
    txDbClient, err := dbClient.Tx(dbCtx)
    if err != nil {
        return err
    }

    err = user.DeleteUser(
              dbCtx,
              txDbClient,
              getUser,
          )
    if err != nil {
        // DBロールバック
        txDbClient.Rollback()

        return err
    }

    // Firebaseユーザー削除
    err = auth.DeleteUser(uid)
    if err != nil {
        // DBロールバック
        txDbClient.Rollback()

        return err
    }

    // DBコミット
    txDbClient.Commit()

    return nil
}
