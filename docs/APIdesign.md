
# REST API仕様

# 作成

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
  "shopping_day": "20200720",
  "product_name": "にんじん",
  "quantity": "1袋",
  "price": "190円"
}
```

### レスポンス
#### 成功時
 - ステータスコード: 200 OK
 
#### レスポンスサンプル
```
{
  "shopping_id": 1,
  "shopping_day": "20200720",
  "product_name": "にんじん",
  "quantity": "1袋",
  "price": "190円"
}
```

#### 失敗時

 - ステータスコード: 400 Bad Request
 - ステータスコード: 409 Conflict
 - ステータスコード: 500 Internal Server Error


# 参照

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
  "shopping_day": "20200720",
  "product_name": "にんじん",
  "quantity": "1袋",
  "price": "190円"
}
{
  "shopping_id": 2,
  "shopping_day": "20200721",
  "product_name": "たまねぎ",
  "quantity": "1袋",
  "price": "210円"
}
{
  "shopping_id": 3,
  "shopping_day": "20200722",
  "product_name": "じゃがいも",
  "quantity": "null",
  "price": "null"
}


```



# 更新


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
  "shopping_day": "20200724"
}
```

### レスポンス
#### 成功時
 - ステータスコード: 200 OK
 
#### レスポンスサンプル

```
{
  "shopping_id": 1,
  "shopping_day": "20200724"
}
```

#### 失敗時

 - ステータスコード: 400 Bad Request
 - ステータスコード: 404 Not Found
 - ステータスコード: 500 Internal Server Error




## リスト内容変更API

### 概要

指定した買い物リストの商品、個数、価格を変更する

### パス

/api/shopping/{shopping_id}/products/{shopping_product_id}

### メソッド

 - PATCH
      - JSON (Req/Res)
 
 
#### リクエストサンプル

```
{
  "product_name": "豚肉",
  "quantity": "100g",
  "price": "240円"
}

```

### レスポンス
#### 成功時
 - ステータスコード: 200 OK
 
#### レスポンスサンプル

```
{
  "shopping_id": 1,
  "product_name": "豚肉",
  "quantity": "100g",
  "price": "240円"
}
```

#### 失敗時

 - ステータスコード: 400 Bad Request
 - ステータスコード: 404 Not Found
 - ステータスコード: 500 Internal Server Error








## リストの商品追加API

### 概要

指定した買い物リストの商品を追加する

### パス

/api/shopping/{shopping_id}/products/{shopping_product_id}

### メソッド

 - POST
      - JSON (Req/Res)
 
 
#### リクエストサンプル

```
{
  "product_name": "たまねぎ"
}

```

### レスポンス
#### 成功時
 - ステータスコード: 200 OK
 
#### レスポンスサンプル

```
{
  "shopping_id": 1,
  "shopping_day": "20200720",
  "product_name": "にんじん",
  "quantity": "1袋",
  "price": "190円"
}
{
  "shopping_id": 1,
  "shopping_day": "20200720",
  "product_name": "たまねぎ"
}
```

#### 失敗時

 - ステータスコード: 400 Bad Request
 - ステータスコード: 404 Not Found
 - ステータスコード: 500 Internal Server Error



## リストの商品削除API

### 概要

指定した買い物リストの商品を削除する

### パス

/api/shopping/{shopping_id}/products/{shopping_product_id}/product_name

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
 - ステータスコード: 500 Internal Server Error




## リストの個数削除API

### 概要

指定した買い物リストの個数を削除する

### パス

/api/shopping/{shopping_id}/products/{shopping_product_id}/quantity

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
 - ステータスコード: 500 Internal Server Error




## リストの価格削除API

### 概要

指定した買い物リストの価格を削除する

### パス

/api/shopping/{shopping_id}/products/{shopping_product_id}/price

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
 - ステータスコード: 500 Internal Server Error







# 削除


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
 - ステータスコード: 500 Internal Server Error
