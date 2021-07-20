
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
 - ステータスコード: 200
 
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


 
