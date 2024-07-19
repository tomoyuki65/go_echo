
> openapi-to-md
> openapi-to-md src/docs/swagger.json

# go_echo API

> Version 1.0

Go言語（Golang）のフレームワーク「Echo」によるバックエンドAPIのサンプル

## Path Table

| Method | Path | Description |
| --- | --- | --- |
| POST | [/hello](#posthello) |  |
| POST | [/post](#postpost) |  |
| GET | [/post/{id}](#getpostid) |  |
| DELETE | [/post/{id}](#deletepostid) |  |
| POST | [/user](#postuser) |  |
| GET | [/user/{uid}](#getuseruid) |  |
| PUT | [/user/{uid}](#putuseruid) |  |
| DELETE | [/user/{uid}](#deleteuseruid) |  |
| GET | [/users](#getusers) |  |

## Reference Table

| Name | Path | Description |
| --- | --- | --- |
| Bearer | [#/components/securitySchemes/Bearer](#componentssecurityschemesbearer) | Type "Bearer" followed by a space and JWT token. |
| api_internal_handlers_post.PostResponse | [#/components/schemas/api_internal_handlers_post.PostResponse](#componentsschemasapi_internal_handlers_postpostresponse) |  |
| api_internal_handlers_user.UserResponse | [#/components/schemas/api_internal_handlers_user.UserResponse](#componentsschemasapi_internal_handlers_useruserresponse) |  |
| index.PostIndexResponse | [#/components/schemas/index.PostIndexResponse](#componentsschemasindexpostindexresponse) |  |
| index.RequestBody | [#/components/schemas/index.RequestBody](#componentsschemasindexrequestbody) |  |
| post.CreatePostRequestBody | [#/components/schemas/post.CreatePostRequestBody](#componentsschemaspostcreatepostrequestbody) |  |
| post.MessageResponse | [#/components/schemas/post.MessageResponse](#componentsschemaspostmessageresponse) |  |
| post.PostResponseWithUser | [#/components/schemas/post.PostResponseWithUser](#componentsschemaspostpostresponsewithuser) |  |
| post.User | [#/components/schemas/post.User](#componentsschemaspostuser) |  |
| user.CreateUserRequestBody | [#/components/schemas/user.CreateUserRequestBody](#componentsschemasusercreateuserrequestbody) |  |
| user.MessageResponse | [#/components/schemas/user.MessageResponse](#componentsschemasusermessageresponse) |  |
| user.Post | [#/components/schemas/user.Post](#componentsschemasuserpost) |  |
| user.UpdateUserRequestBody | [#/components/schemas/user.UpdateUserRequestBody](#componentsschemasuserupdateuserrequestbody) |  |
| user.UserResponseWithPosts | [#/components/schemas/user.UserResponseWithPosts](#componentsschemasuseruserresponsewithposts) |  |

## Path Details

***

### [POST]/hello

- Description  
パラメータのtextを出力する

#### RequestBody

- application/json

```ts
{
  text: string
}
```

#### Responses

- 200 出力メッセージ

`*/*`

```ts
{
  message?: string
}
```

- 400 Bad Request

- 500 Internal Server Error

***

### [POST]/post

- Description  
Post作成

#### RequestBody

- application/json

```ts
{
  text: string
  user_id: integer
}
```

#### Responses

- 201 Created

`*/*`

```ts
{
  created_at?: string
  edges: {
  }
  id?: integer
  text?: string
  updated_at?: string
  user_id?: integer
}
```

- 400 Bad Request

- 500 Internal Server Error

***

### [GET]/post/{id}

- Description  
有効な対象Post取得

#### Responses

- 200 OK

`*/*`

```ts
{
  created_at?: string
  edges: {
    users: {
      created_at?: string
      edges: {
      }
      email?: string
      first_name?: string
      id?: integer
      last_name?: string
      uid?: string
      updated_at?: string
    }
  }
  id?: integer
  text?: string
  updated_at?: string
  user_id?: integer
}
```

- 404 Not Found

- 405 Method Not Allowed

- 500 Internal Server Error

***

### [DELETE]/post/{id}

- Description  
対象Post削除

#### Responses

- 200 OK

`*/*`

```ts
{
  message?: string
}
```

- 404 Not Found

- 405 Method Not Allowed

- 500 Internal Server Error

***

### [POST]/user

- Description  
ユーザー作成

#### RequestBody

- application/json

```ts
{
  email: string
  first_name: string
  last_name: string
  uid: string
}
```

#### Responses

- 201 Created

`*/*`

```ts
{
  created_at?: string
  edges: {
  }
  email?: string
  first_name?: string
  id?: integer
  last_name?: string
  uid?: string
  updated_at?: string
}
```

- 400 Bad Request

- 500 Internal Server Error

***

### [GET]/user/{uid}

- Description  
有効な対象ユーザー取得

- Security  
Bearer  

#### Responses

- 200 OK

`*/*`

```ts
{
  created_at?: string
  edges: {
    posts: {
      created_at?: string
      edges: {
      }
      id?: integer
      text?: string
      updated_at?: string
      user_id?: integer
    }[]
  }
  email?: string
  first_name?: string
  id?: integer
  last_name?: string
  uid?: string
  updated_at?: string
}
```

- 401 Unauthorized

- 404 Not Found

- 405 Method Not Allowed

- 500 Internal Server Error

***

### [PUT]/user/{uid}

- Description  
対象ユーザー更新

#### RequestBody

- application/json

```ts
{
  email?: string
  first_name?: string
  last_name?: string
  password?: string
}
```

#### Responses

- 200 OK

`*/*`

```ts
{
  created_at?: string
  edges: {
  }
  email?: string
  first_name?: string
  id?: integer
  last_name?: string
  uid?: string
  updated_at?: string
}
```

- 404 Not Found

- 405 Method Not Allowed

- 500 Internal Server Error

***

### [DELETE]/user/{uid}

- Description  
対象ユーザー削除

#### Responses

- 200 OK

`*/*`

```ts
{
  message?: string
}
```

- 404 Not Found

- 405 Method Not Allowed

- 500 Internal Server Error

***

### [GET]/users

- Description  
有効な全てのユーザー取得

#### Responses

- 200 OK

`*/*`

```ts
{
  created_at?: string
  edges: {
  }
  email?: string
  first_name?: string
  id?: integer
  last_name?: string
  uid?: string
  updated_at?: string
}[]
```

- 500 Internal Server Error

## References

### #/components/securitySchemes/Bearer

```ts
{
  "description": "Type \"Bearer\" followed by a space and JWT token.",
  "type": "apiKey",
  "name": "Authorization",
  "in": "header"
}
```

### #/components/schemas/api_internal_handlers_post.PostResponse

```ts
{
  created_at?: string
  edges: {
  }
  id?: integer
  text?: string
  updated_at?: string
  user_id?: integer
}
```

### #/components/schemas/api_internal_handlers_user.UserResponse

```ts
{
  created_at?: string
  edges: {
  }
  email?: string
  first_name?: string
  id?: integer
  last_name?: string
  uid?: string
  updated_at?: string
}
```

### #/components/schemas/index.PostIndexResponse

```ts
{
  message?: string
}
```

### #/components/schemas/index.RequestBody

```ts
{
  text: string
}
```

### #/components/schemas/post.CreatePostRequestBody

```ts
{
  text: string
  user_id: integer
}
```

### #/components/schemas/post.MessageResponse

```ts
{
  message?: string
}
```

### #/components/schemas/post.PostResponseWithUser

```ts
{
  created_at?: string
  edges: {
    users: {
      created_at?: string
      edges: {
      }
      email?: string
      first_name?: string
      id?: integer
      last_name?: string
      uid?: string
      updated_at?: string
    }
  }
  id?: integer
  text?: string
  updated_at?: string
  user_id?: integer
}
```

### #/components/schemas/post.User

```ts
{
  created_at?: string
  edges: {
  }
  email?: string
  first_name?: string
  id?: integer
  last_name?: string
  uid?: string
  updated_at?: string
}
```

### #/components/schemas/user.CreateUserRequestBody

```ts
{
  email: string
  first_name: string
  last_name: string
  uid: string
}
```

### #/components/schemas/user.MessageResponse

```ts
{
  message?: string
}
```

### #/components/schemas/user.Post

```ts
{
  created_at?: string
  edges: {
  }
  id?: integer
  text?: string
  updated_at?: string
  user_id?: integer
}
```

### #/components/schemas/user.UpdateUserRequestBody

```ts
{
  email?: string
  first_name?: string
  last_name?: string
  password?: string
}
```

### #/components/schemas/user.UserResponseWithPosts

```ts
{
  created_at?: string
  edges: {
    posts: {
      created_at?: string
      edges: {
      }
      id?: integer
      text?: string
      updated_at?: string
      user_id?: integer
    }[]
  }
  email?: string
  first_name?: string
  id?: integer
  last_name?: string
  uid?: string
  updated_at?: string
}
```
