# bit-board-sendData

## 基本仕様
---
### 概要
BitBoardのデータを送信する機能

### 特徴
必要なデータを取得,整形する

## 機能仕様
---
### 利用法

- req
```cassandraql
{
    "userID":"sample",
    "password":"sample"
}
```


## 詳細仕様
deploy Lambda Name：
deploy API Gateway：

---
### アーキテクチャ


### 依存環境
- AWS (Lambda)
- AWS (API Gateway)

### 使用ライブラリ
主要なライブラリのみを示す
- aws-sdk()
- Gin(Webフレームワーク)

### デプロイ
AWS Lambda：SendSignageDate

### セットアップ
対象となるソースコードをzipに圧縮してアップロードする必要があります
2. GOOS=linux GOARCH=amd64 go build -o hello main.go
3. zip function.zip hello

### 注意点
動作確認はAWSのLambdaにアップロードする必要があります
zipはGitにあげないこと
OpenWeatherのAPIを使用しているためリクエスト数などに注意

### 作成者
Taketo Wakamatsu (若松丈人)
