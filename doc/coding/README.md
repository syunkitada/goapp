# Coding Rule

## パッケージ

- main パッケージは、cmd/ 配下に格納される
- main 以外のパッケージは、pkg/ 配下に格納される
- ファイル名は必ず lowercase を使う
- パッケージ名から中身が想像できるようにする
  - また、パッケージコメント(package 宣言の直上コメント)を必ず書くこと
- 衝突の恐れがある名前は避ける
  - snake_case は使ってもよい（一般的にはダメと言われてるが、無理に単語をつなげるより見やすい)
  - NG
    - pkg/resource/api
    - pkg/resource/controller
    - pkg/resourceapi
    - pkg/resourcecontroller
  - OK
    - pkg/resource/resource_api
    - pkg/resource/resource_controller
- pkg/ 直下は極力汚さないようにする
  - 各サービスを示す大枠の namespace を切って、その中で機能を拡張していく
    - pkg/authproxy/...
    - pkg/resource/...
  - 各サービス共通で使うようなものは pkg/lib に含める
    - pkg/lib/logger
    - pkg/lib/config
    - pkg/lib/base_app
    - pkg/lib/base_client
  - logger, config など衝突の恐れがあるが、各サービス共通で使うものなのでセーフ
    - 外部の logger や config を別途利用する場合には注意が必要
- type 名や変数名は省略文字も含めてすべて CamelCase を利用する
  - lowercase と CamelCase を自動変換する場合に、省略語かどうかを考慮させないため
  - HTTPClient ではなく HttpClient とする
  - ただし、外部ライブラリの利用時は例外とする

## ファイルの整形、Lint

- ファイルの整形には goimports を利用する
- 静的解析ツールには gometalinter を利用する
  - vet, golint, errcheck を有効
