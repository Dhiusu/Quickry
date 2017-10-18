# Quickly
簡単にスタブ環境を用意できるスタブアプリです。
返却するレスポンス情報は、config.jsonで設定できます。

## 手順
- go run main.go

## レスポンス変更
- ./config.json
---
`{
  "routes":[
    {
      "url":"/abc/def/aaa",
      "response":"{\"response\":[aaa]}}"
    },
    ...
  ]
}`


