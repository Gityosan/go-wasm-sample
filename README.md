# go-wasm-sample

フロントエンドの重たい処理をGoで軽く処理させるやり方を学ぶためのリポジトリです。

#### 起動方法
```bash 
tinygo build -o public/main.wasm -target wasm src/main.go
go run server.go
```