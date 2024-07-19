package post

import (
    "context"
    "fmt"

    "api/ent"
    "api/internal/repositories/post"
)

type PostService interface {
    CreatePost(
        echoCtx context.Context,
        dbCtx context.Context,
        dbClient *ent.Client,
        userId int64,
        text string,
    ) (*ent.Post, error)
    GetPost(
        echoCtx context.Context,
        dbCtx context.Context,
        dbClient *ent.Client,
        id int64,
    ) (PostResponse, error)
    DeletePost(
        echoCtx context.Context,
        dbCtx context.Context,
        dbClient *ent.Client,
        id int64,
    ) error
}

type postService struct{}

func NewPostService() PostService {
    return &postService{}
}

// レスポンス定義
type Post struct { *ent.Post }
type EmptyResponse map[string]interface{}
type PostResponse interface {}
func (u Post) PostResponse() {}
func (e EmptyResponse) PostResponse() {}

func (s *postService) CreatePost(
    echoCtx context.Context,
    dbCtx context.Context,
    dbClient *ent.Client,
    userId int64,
    text string,
) (*ent.Post, error) {

    createPost, err := post.CreatePost(
                           dbCtx,
                           dbClient,
                           userId,
                           text,
                       )
    if err != nil {
        return nil, err
    }

    return createPost, nil
}

func (s *postService) GetPost(
    echoCtx context.Context,
    dbCtx context.Context,
    dbClient *ent.Client,
    id int64,
) (PostResponse, error) {

    getPost, err := post.GetPostFromId(
                        dbCtx,
                        dbClient,
                        id,
                    )
    if err != nil {
        return nil, err
    }

    if getPost == nil {
        emptyResponse := map[string]interface{}{}
        return emptyResponse, nil
    }

    return getPost, nil
}

func (s *postService) DeletePost(
    echoCtx context.Context,
    dbCtx context.Context,
    dbClient *ent.Client,
    id int64,
) error {

    getPost, err := post.GetPostFromId(
                        dbCtx,
                        dbClient,
                        id,
                    )
    if err != nil {
        return err
    }

    if getPost == nil {
        return fmt.Errorf("no post")
    }

    err = post.DeletePost(
              dbCtx,
              dbClient,
              getPost,
          )
    if err != nil {
        return err
    }

    return nil
}