
> openapi-to-md
> openapi-to-md src/docs/swagger.json

# go_echo API

> Version 1.0

Go言語（Golang）のフレームワーク「Echo」によるバックエンドAPIのサンプル

## Path Table

| Method | Path | Description |
| --- | --- | --- |
| POST | [/hello](#posthello) |  |

## Reference Table

| Name | Path | Description |
| --- | --- | --- |
| index.PostIndexResponse | [#/components/schemas/index.PostIndexResponse](#componentsschemasindexpostindexresponse) |  |
| index.RequestBody | [#/components/schemas/index.RequestBody](#componentsschemasindexrequestbody) |  |

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

## References

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
