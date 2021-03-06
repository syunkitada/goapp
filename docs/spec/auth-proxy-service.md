# AuthProxy Service

## Overview

- AuthProxy
  - API Gateway
    - クライアントの認証・認可を行い、各種サービスにプロキシするサービス
    - すべてのリクエストはこのプロキシを経由する
  - クライアントは、初回に必ずユーザ名、パスワードによりトークンを取得し、そのトークンによりサービスを利用する
    - 「ユーザ名・パスワード」により認証し、JWT トークンを発行する
      - トークンには以下の情報を含む
        - ユーザ名
        - トークン有効期限
    - 認証もとによってレスポンスの形式を分ける
      - CLI 用認証
        - トークン
        - 所属プロジェクト一覧
      - ブラウザ用認証
        - トークンは Cookie にセットして返す
          - Secure フラグ, HttpOnly フラグをセットする
        - 所属プロジェクト一覧
        - 初期ステート情報
    - 初回認証時には利用可能なサービス一覧を取得でき、サービス名によりサービスを利用する
  - クライアントは、リクエスト時にトークン、サービス名、メソッド名を指定し、適切なデータをセットしてリクエストする
    - ブラウザからの場合は、トークンは Cookie にセットされてリクエストされる
  - AuthProxy は、トークンとサービス名とメソッド名によって認証・認可を行い、ルーティングする

## Data Model

- User
  - ユーザデータそのもの
  - ユーザ名、パスワードを保持
- Role
  - User は複数の Project に属することができる
  - Project に属するとき、"admin"や"member"などの Role を必ず一つだけ紐図けられる
  - この Role により、Project 内での User 権限を制御する
    - デフォルト Role
      - "admin"
        - Project の編集や User の追加削除ができる
      - "member"
        - 各種サービスの基本操作
- Project
  - プロジェクトデータそのもの
- ProjectRole
  - プロジェクトに"admin"、"system"、"tenant"などの ProjectRole を必ず一つだけ紐図けられる
  - この ProjectRole により、クラスタ内での Project 権限をする
    - デフォルト ProjectRole
      - "admin"
        - 管理者権限
        - すべてを操作が可能
      - "system"
        - システム権限
        - システム上必要なすべての操作が可能
      - "tenant"
        - 一般ユーザテナント権限
        - テナント内に閉じた操作が可能

## 認証システム

### JWT トークン

- JWT トークンは、署名付きのデータを保持したトークンで、Header、Payload、Signature の 3 つからなる
  - xxxx[Header].yyyy[Payload].zzzz[Signature]
- Header: 以下の情報を Base64UrlEncode したもの
  - alg: Signature を生成するために利用するアルゴリズム
  - typ: トークンのタイプ(通常は JWT)
- Payload: 任意の情報を Base64UrlEncode したもの
  - Token の有効期限や、公開してもよいユーザのセッション情報などを保存する
- Signature: Header と Payload に Secret(秘密文字列)をつなげてハッシュ化したもの
  - Header と Payload は誰でも中身を見ることができ、Secret を知る者(発行者など)だけがトークンが偽装されたかどうか検証できる

### JWT トークン利用時の留意事項

- Web ブラウザから利用する場合、JWT トークンは Cookie に保存する
  - Secure フラグをセットする
    - Cookie は HTTPS でのみ扱われる
    - HTTP では、Cookie の中身が中間者に見えてしまうため、トークンが漏洩する可能性がある
  - HTTPOnly フラグをセットする
    - Cookie はクライアントスクリプト(JavaScript など)から読めなくなる
  - 仮に Web ブラウザのローカルストレージで JWT トークンを管理した場合
    - これは JavaScript から読めてしまうため、セキュアではない
    - 悪意ある JavaScript コードが混入した場合、トークンが漏洩してしまう
- トークンは純粋なセッション管理のみに利用する
  - トークンが保証するのは、Web サーバがそのユーザ宛てに発行したものであることのみ
    - これが保証されるため、サーバ側でセッション用のデータベースを持つ必要がない
    - トークンはユーザ存在やそのユーザの権限を保証できるものではない
      - トークンでユーザ存在や権限を管理してしまうと、(権限取り消し前の)古いトークンでもそのユーザのリソースを利用できてしまう
  - トークンつきのリクエストが来た場合は、そのトークンのユーザ名からデータベースを引き当てて、そのユーザの存在チェック、権限チェックを行う

## サービス概念

- サービスはユーザごとに提供されるものプロジェクトごとに提供されるものに分かれる
  - ユーザ向けサービス
    - Dashboard, Wiki, Chat など
  - プロジェクト向けサービス
    - Resource, Monitor など
- トークン取得時に、トークンとは別にユーザの利用可能なプロジェクト一覧とサービス一覧を返し、ユーザはそれに基づいて任意の操作を行う
- サービスは ProjectRole によって利用可能かどうかを判定する

## アクション認可

- アクションはサービスに属し、プロジェクトロール単位、ロール単位で認可される
- アクションにはプロジェクトロールが必ず指定され、まずプロジェクトロールでバリデートし、加えてアクションにロールが指定されてる場合はロールでもバリデートする

## プロキシ方法

- AuthProxy とサービス間も認証および TLS で暗号化されてる必要がある
- サービスは発行されたサービスアカウントによって AuthProxy で認証を行い、自身のサービスを AuthProxy に登録する
- サービス登録にはエンドポイントの一覧、API 一覧、UI 情報、認可情報、トークンを AuthProxy に登録し、AuthProxy はこのトークンによりサービスへアクセスする
  - サービス側も AuthProxy のための簡易的なトークン認証を持つ
- サービスはトークンの有効期限が切れる前に、定期的に自身のサービス情報を更新する必要がある
- 一定期間更新のないサービスは利用できなくなる

## ロードバランシング

- ユーザは、ドメイン名によりアクセスしてくる
- 各地域に authproxy を配置する場合は、GDNS によりユーザ IP から地域を特定して、近い authproxy の VIP へ誘導する
- authproxy は、それごとに各サービスのエンドポイント一覧とレイテンシ情報とステータス情報をもっており優先度を決定する
  - authproxy は、inmemory で エンドポイント一覧のレイテンシ情報とステータス情報を定期的に更新する
  - 特定の エンドポイント の負荷が偏った場合、レイテンシが悪化し、優先度が下がる
  - 特定の エンドポイント がメンテナンスや障害によりダウンした場合、この定期更新によりダウンを検知し利用されなくなる
- 優先度が近似してるものが複数ある場合、ソースハッシュで並び替えて、順にアクセスしていく

## クラスタとノード

- 各ノードは、自身の所属クラスタのみを意識する
- 各クラスタはクラスタ API とクラスタコントローラを持つ
  - クラスタ API 群がサービスとなる
- 各クラスタ API は、クラスタ DB 、上位クラスタ、下位クラスタのみを意識する
  - クラスタは、ツリー構造として拡張でき、上位クラスタと下位クラスタを持つことができる
  - クラスタへのリクエストは、DB のみで完結させる
  - 完結できない場合は、タスクを DB に書き込み、コントローラが非同期でタスクを処理する
- 上位ノード、下位ノード間の連携
  - 原則、ノード間通信は上位ノードと下位ノードでのみ通信を行う
    - ノードから見て、自身は下位ノード、クラスタ API が上位ノード
    - クラスタ API からみて、自身は下位ノード、上位クラスタのクラスタ API が上位ノード
  - 上位ノード は、Authproxy と同等の認証機能を持つ
    - 各ノードはユーザ、パスワードにより、中央ノードにアクセスしてトークンを取得し、トークンにより認証、リクエストを行う
    - ノードのデプロイ時に、ユーザ、パスワードの設定を配る
  - 各ノードは、定期的にローカルにトークンを生成し、数世代分保持しており、下位ノードへの認証はこのトークンにより行うことができる
  - ノードは定期的にクラスタ API にレポートを行い、自身のトークンを上位ノードに登録し、自身が実行すべきタスクを取得して処理する

## UI

- クライアントの WebUI や CUI は、 AuthProxy で認証し、UI 情報を取得して、UI を自動生成する
- WebUI の認証の流れ
  - Cookie のトークンによって、LoginWithToken を実行し、認可情報を取得する
  - LoginWithToken に失敗する場合は、ログイン画面を表示する
    - ログイン画面から、ユーザ、パスワード認証を行い、トークンと認可情報を取得する
      - このとき Cookie にトークンを保存する
  - 認可情報には、プロジェクト一覧とサービス一覧を含んでおり、これにより UI のフレームを生成する
  - User サービスにアクセスしようとしてる場合は、GetServiceDashboardIndex によって Index を取得する
  - Project サービスにアクセスしようとしてる場合は、GetProjectServiceDashboardIndex によって Index を取得する
  - Index 情報からサービス UI を自動生成する
- CUI の認証の流れ
  - ユーザ、パスワード認証を行い、トークンと認可情報を取得する
  - User サービスにアクセスしようとしてる場合は、GetServiceIndex によって Index を取得する
  - Project サービスにアクセスしようとしてる場合は、GetProjectServiceIndex によって Index を取得する
  - Index 情報に利用できるコマンド情報が含まれている

## ACL

- ブラックボックス or ホワイトボックスで IP の制限を設定できるようにする
- DOS 対策
- TODO

## キャッシング

- TOD
