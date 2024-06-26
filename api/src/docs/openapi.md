
> openapi-to-md
> openapi-to-md src/docs/swagger.json

# go_echo API

> Version 1.0

Go言語（Golang）のフレームワーク「Echo」によるバックエンドAPIのサンプル

## Path Table

| Method | Path | Description |
| --- | --- | --- |
| POST | [/hello](#posthello) |  |
| POST | [/user](#postuser) |  |
| GET | [/user/:uid](#getuseruid) |  |
| PUT | [/user/:uid](#putuseruid) |  |
| DELETE | [/user/:uid](#deleteuseruid) |  |
| GET | [/users](#getusers) |  |

## Reference Table

| Name | Path | Description |
| --- | --- | --- |
| api_internal_handlers_user.UserResponse | [#/components/schemas/api_internal_handlers_user.UserResponse](#componentsschemasapi_internal_handlers_useruserresponse) |  |
| index.PostIndexResponse | [#/components/schemas/index.PostIndexResponse](#componentsschemasindexpostindexresponse) |  |
| index.RequestBody | [#/components/schemas/index.RequestBody](#componentsschemasindexrequestbody) |  |
| user.CreateUserRequestBody | [#/components/schemas/user.CreateUserRequestBody](#componentsschemasusercreateuserrequestbody) |  |
| user.MessageResponse | [#/components/schemas/user.MessageResponse](#componentsschemasusermessageresponse) |  |
| user.UpdateUserRequestBody | [#/components/schemas/user.UpdateUserRequestBody](#componentsschemasuserupdateuserrequestbody) |  |

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

### [GET]/user/:uid

- Description  
有効な対象ユーザー取得

#### Responses

- 200 OK

`*/*`

```ts
{
  created_at?: string
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

### [PUT]/user/:uid

- Description  
対象ユーザー更新

#### RequestBody

- application/json

```ts
{
  email?: string
  first_name?: string
  last_name?: string
}
```

#### Responses

- 200 OK

`*/*`

```ts
{
  created_at?: string
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

### [DELETE]/user/:uid

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

### #/components/schemas/api_internal_handlers_user.UserResponse

```ts
{
  created_at?: string
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

### #/components/schemas/user.UpdateUserRequestBody

```ts
{
  email?: string
  first_name?: string
  last_name?: string
}
```
