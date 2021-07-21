
# REST API仕様

## 買い物リスト作成API

### 概要

買い物リストを作成する

### パス

/api/shopping

### メソッド

 - POST
     - JSON (Req/Res)
 
 
#### リクエストサンプル
```
{
  "shopping_day": "2020-07-20",
  "product_name": "にんじん",
  "quantity": "1袋",
  "price": "190円",
}
```

### レスポンス
#### 成功時
 - ステータスコード: 200 OK
 
#### レスポンスサンプル
```
{
  "shopping_id": 1,
  "shopping_day": "2020-07-20",
  "product_name": "にんじん",
  "quantity": "1袋",
  "price": "190円",
}
```

#### 失敗時

 - ステータスコード: 400 Bad Request






## 参照API

### 概要

全買い物リストを参照する

### パス

/api/shopping

### メソッド

 - GET
     - JSON (Res)
 
 
#### リクエストサンプル

空

### レスポンス
#### 成功時
 - ステータスコード: 200 OK
 
#### レスポンスサンプル
```
{
  "shopping_id": 1,
  "shopping_day": "2020-07-20",
  "product_name": "にんじん",
  "quantity": "1袋",
  "price": "190円",
}
{
  "shopping_id": 2,
  "shopping_day": "2020-07-21",
  "product_name": "たまねぎ",
  "quantity": "1袋",
  "price": "210円",
}
{
  "shopping_id": 3,
  "shopping_day": "2020-07-22",
  "product_name": "じゃがいも",
  "quantity": "null",
  "price": "null",
}


```

#### 失敗時

 - ステータスコード: 400 Bad Request






## 日付更新API

### 概要

指定した買い物リストの日付を変更する

### パス

/api/shopping/{shopping_id}

### メソッド

 - PATCH
      - JSON (Req/Res)
 
 
#### リクエストサンプル

```
{
  "shopping_day": "2020-07-24",
}
```

### レスポンス
#### 成功時
 - ステータスコード: 200 OK
 
#### レスポンスサンプル

```
{
  "shopping_id": 1,
  "shopping_day": "2020-07-24",
}
```

#### 失敗時

 - ステータスコード: 400 Bad Request
 - ステータスコード: 404 Not Found















## 削除API

### 概要

指定した買い物リストを削除する

### パス

/api/shopping/{shopping_id}

### メソッド

 - DELETE
 
 
#### リクエストサンプル

空

### レスポンス
#### 成功時
 - ステータスコード: 204 No Content
 
#### レスポンスサンプル

空

#### 失敗時

 - ステータスコード: 400 Bad Request
 - ステータスコード: 404 Not Found
 
