# auth-with-golang-vue3
Golang + Vue.js + docker を用いたアプリケーションのサンプル

# バージョン
Go => 1.15

# 起動方法
```
docker-compose up
```


# APIの基本アクセス
## ユーザー登録
### エンドポイント
http://localhost/api/register

### メソッド
POST

### リクエストデータ
```
{
  "first_name": "michi",
  "last_name": "kaneko",
  "email": "test@example.com",
  "password": "test",
  "password_confirm": "test"
}
```

## ログイン
### エンドポイント
http://localhost/api/login

### メソッド
POST

### リクエストデータ
```
{
  "email": "test@example.com",
  "password": "test",
}
```

## ユーザーデータの表示
### エンドポイント
http://localhost/api/user

### メソッド
GET

## ログアウト
### エンドポイント
http://localhost/api/logout

### メソッド
GET

## パスワードリセットメール送信
### エンドポイント
http://localhost/api/forgot

### メソッド
POST

### リクエストデータ
```
{
  "email": "test@example.com",
}
```

## パスワード再設定
### エンドポイント
http://localhost/api/reset

### メソッド
POST

### リクエストデータ
```

{
  "token": "メールに記載されたトークン ",
  "password": "abc",
  "password_confirm": "abc"
}
```

