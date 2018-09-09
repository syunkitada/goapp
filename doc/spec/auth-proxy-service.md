# AuthProxy Service


## Overview
* AuthProxy
    * クライアントの認証・認可を行い、各種サービスにプロキシするサービス
        * すべてのリクエストはこのプロキシを経由する
    * クライアントは、初回に必ずトークン認証をする
        * 「ユーザ名・パスワード・プロジェクト名」により認証し、JWTトークンを発行する
        * トークンには以下の情報を含む
            * ユーザ名
            * トークン有効期限
    * クライアントは、リクエスト時にトークンと、サービス名、メソッド名を指定し、適切なデータをセットしてリクエストする
    * プロキシは、トークンとサービス名とメソッド名によって認証・認可を行い、バックエンドサービスに適切にルーティングする
        * サービスのルーティング、メソッドの認可の設定は、プロキシ起動時にservice.dディレクトリに保存したyamlで行う
    * プロキシとバックエンドサービスとのやり取りは、gRPCによって行う


## Data Model
* User
    * ユーザデータそのもの
    * ユーザ名、パスワードを保持
* Role
    * Userは複数のProjectに属することができる
    * Projectに属するとき、"admin"や"member"などのRoleによってUserとProjectは紐図けられる
    * このRoleにより、Project内でのUser権限を制御する
        * デフォルトRole
            * "admin"
                * Projectの編集やUserの追加削除ができる
            * "member"
                * 各種サービスの基本操作
* Project
    * プロジェクトデータそのもの
* ProjectRole
    * プロジェクトに"admin"、"system"、"tenant"などのProjectRoleを紐図けられる
    * このProjectRoleにより、クラスタ内でのProject権限をする
        * デフォルトProjectRole
            * "admin"
                * 管理者権限
                * すべてを操作が可能
            * "system"
                * システム権限
                * システム上必要なすべての操作が可能
            * "tenant"
                * 一般ユーザテナント権限
                * テナント内に閉じた操作が可能


## 認証システム

### JWTトークン
* JWTトークンは、署名付きのデータを保持したトークンで、Header、Payload、Signatureの3つからなる
    * xxxx[Header].yyyy[Payload].zzzz[Signature]
* Header: 以下の情報をBase64UrlEncodeしたもの
    * alg: Signatureを生成するために利用するアルゴリズム
    * typ: トークンのタイプ(通常はJWT)
* Payload: 任意の情報をBase64UrlEncodeしたもの
    * Tokenの有効期限や、公開してもよいユーザのセッション情報などを保存する
* Signature: HeaderとPayloadにSecret(秘密文字列)をつなげてハッシュ化したもの
    * HeaderとPayloadは誰でも中身を見ることができ、Secretを知る者(発行者など)だけがトークンが偽装されたかどうか検証できる


### JWTトークン利用時の留意事項
* Webブラウザから利用する場合、JWTトークンはCookieに保存する
    * Webブラウザなどのローカルストレージで管理する場合、セキュアではなくJavaScriptから読めてしまう
        * 悪意あるJavaScriptコードが混入してもJWTトークンは見れないため、トークン漏洩の対策となる
    * Cookieの場合、設定によってJavascriptから触ることはできない
* トークンは純粋なセッション管理のみに利用する
    * トークン情報にユーザ存在チェックや権限チェックのための情報を混ぜない
        * これらの情報を混ぜてしまうと、(権限取り消し前の)古いトークンでもリソースにアクセスできてしまう
