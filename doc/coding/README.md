# Coding Rule


## パッケージ
* mainパッケージは、cmd/ 配下に格納される
* main以外のパッケージは、pkg/ 配下に格納される
* ファイル名は必ずlowercaseを使う
* パッケージ名から中身が想像できるようにする
    * また、パッケージコメント(package宣言の直上コメント)を必ず書くこと
* 衝突の恐れがある名前は避ける
    * snake_caseは使ってもよい（一般的にはダメと言われてるが、無理に単語をつなげるより見やすい)
    * NG
        * pkg/resource/api
        * pkg/resource/controller
        * pkg/resourceapi
        * pkg/resourcecontroller
    * OK
        * pkg/resource/resource_api
        * pkg/resource/resource_controller
* pkg/ 直下は極力汚さないようにする
    * 各サービスを示す大枠のnamespaceを切って、その中で機能を拡張していく
        * pkg/authproxy/...
        * pkg/resource/...
    * 各サービス共通で使うようなものは pkg/libに含める
        * pkg/lib/logger
        * pkg/lib/config
        * pkg/lib/base_app
        * pkg/lib/base_client
    * logger, configなど衝突の恐れがあるが、各サービス共通で使うものなのでセーフ
        * 外部のloggerやconfigを別途利用する場合には注意が必要


## ファイルの整形、Lint
* ファイルの整形にはgoimportsを利用する
* 静的解析ツールにはgometalinterを利用する
    * vet, golint, errcheckを有効
