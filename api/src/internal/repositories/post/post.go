package post

import (
    "context"
    "fmt"
    "time"

    "api/ent"
    entPost "api/ent/post"
    entUser "api/ent/user"
)

func CreatePost(
    dbCtx context.Context,
    dbClient *ent.Client,
    userId int64,
    text string,
) (*ent.Post, error) {

    post, err := dbClient.Post.
                          Create().
                          SetUserID(userId).
                          SetText(text).
                          Save(dbCtx)
    if err != nil {
        return nil, fmt.Errorf("failed creating post: %v", err)
    }

    return post, nil
}

func GetPostFromId(
    dbCtx context.Context,
    dbClient *ent.Client,
    id int64,
) (*ent.Post, error) {

    post, err := dbClient.Post.
                          Query().
                          Where(entPost.IDEQ(id)).
                          Where(entPost.DeletedAtIsNil()).
                          WithUsers(func(query *ent.UserQuery) {
                              query.Where(entUser.DeletedAtIsNil())
                          }).
                          First(dbCtx)
    if err != nil && !ent.IsNotFound(err) {
        return nil, fmt.Errorf("failed get post: %v", err)
    }

    return post, nil
}

func DeletePost(
    dbCtx context.Context,
    dbClient *ent.Client,
    post *ent.Post,
) error {

    _, err := dbClient.Post.
                       UpdateOne(post).
                       SetDeletedAt(time.Now()).
                       Save(dbCtx)
    if err != nil {
        return fmt.Errorf("failed delete post: %v", err)
    }

    return nil
}