# GoのEchoによるAPI開発のサンプル  
Go言語（Golang）のフレームワーク「Echo」による  
バックエンドAPI開発のサンプルです。  
  
## 機能一覧  
・ユーザー作成（POST /user）  
・有効な対象ユーザー取得（GET /user/{uid}）  
・ユーザー更新（PUT /user/{uid}）  
・ユーザー論理削除（DELETE /user/{uid}）  
・有効な全てのユーザー取得（GET /users）  
  
## 使用技術  
Go "1.22.3"  
Echo "v4.12.0"  
entgo.io/ent "v0.13.1"  
atlas "v0.24.1-764acbb-canary"  
Docker  
docker-compose  
Firebase Authentication  
  
## 参考記事  
[・Go言語（Golang）のEchoでバックエンドAPIの開発方法まとめ](https://golang.tomoyuki65.com/how-to-develop-an-api-with-echo-in-go)  
  