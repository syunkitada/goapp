# Code Generator

- 各サービスの API の仕様は、spec/spec.go によって管理されており、spec/genpkg に以下のファイルを自動生成して利用されています
  - api.go
    - API のインターフェイス
    - この API インターフェイスにしたがって API の実装を行います
  - client.go
    - クライアント用のコード
  - cmd.go
    - CUI 用のコード
  - view.go
    - GUI 用のコード
- spec/genpkg は、以下のコマンドにより自動生成されます。

```
$ make gen

$ ls pkg/resource/resource_api/spec/genpkg
api.go  client.go  cmd.go  view.go
```
