# フォルダ構成
[こちら](https://qiita.com/sueken/items/87093e5941bfbc09bea8)を参考にした
- /cmd: アプリケーションのエントリーポイントとなるファイルを配置する
- /internal: /cmdからしか使わないライブラリを配置する
- /pkg: どのアプリケーションから利用されてもいいようなライブラリを配置する
- /config: 設定ファイルを配置する
- /script: スクリプトを配置する
- /test: テストファイルを配置する
- /docs: ドキュメントを配置する


# ローカルでの起動方法

- マイグレートの行をコメントアウトする
- 以下のコマンドを実行する
```
$ docker compose up -d --build
```


# ドキュメント更新
```
swag init --parseDependency --parseInternal -g ./cmd/api/main.go
```
