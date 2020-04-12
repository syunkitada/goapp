# Spec

## Design

#### Containers (src/containers)

- すべての根幹をつかさどる Component 群
- Root.tsx
  - Root の Component
- AuthRoute.tsx
  - 認証を管理するための Component
  - 認証済みかどうかをチェックし、認証済みでないなら、ログイン用の URL へリダイレクトする

#### Components (src/components)

- UI 表示のための Components 群
- StatelessComponent
  - State を管理しない Component
  - Function で実装する
    - 再レンダリングのコストをなくせる
- StateComponent
  - 更新されたステート情報に基づいて愚直に UI を生成する
  - 基本的にステートは、redux の connect を利用して更新する
  - Index に DataKey を一つ持ち、ステートから Data を取得し、Data が更新されたかで、再レンダリングを行う
    - shouldComponentUpdate()を正しく実装する
- frames/Dashboard.tsx
  - UI の一番外側のフレームを管理するための Component
  - renderIndex で生成した Root の Component を一つ持ち、表示する

#### Apps (src/apps)

- Action、Reducer の定義、処理を行うためのアプリケーション群
- 基本的に Action は Component によって発行され、処理され、Reducer によってステートが更新される

#### Lib (src/lib)

- ライブラリ
- 汎用的な関数などを管理する

## State の扱いについて

- 方針
  - Component は、State の変化に伴って、レンダリングを行う
  - State は、Global と Private を意識する
    - Private は、その Component で完結するもの
    - Grobal は、その Component 以外でも参照されるもの
  - GlobalState の更新はすべて Action および、Reducer によって更新する
    - 画面遷移時などは、Action(actions.service.serviceGetQueries)を実行して、ステートを更新する
    - 遷移先は、2 回レンダリングされる
      - Loading: ローディング画面をレンダリングする
      - Activate: データを非同期で取得したらその結果を用いてレンダリングする
  - GlobalState の情報は、基本的に子の Component に伝搬させてはならない
    - 再レンダリングのコストを許容できる場合は、伝搬させて良い
- 基本的なステートの更新タイミングは以下
  - 初回ロード時(認証直後)
  - Dashboard のプロジェクト切り替え時
  - Dashboard のサービス切り替え時
  - Components の特定イベント時
- Components による更新タイミングは以下
  - panels/ExpansionPanels
    - componentWillMount
      - URL から getQueries を実行する
      - actions.service.serviceGetQueries
    - handleChange
      - URL の変化がないので、うまくやる
      - actions.service.serviceGetQueries
  - panes/Panes
    - componentWillMount
      - URL から getQueries を実行する
      - actions.service.serviceGetQueries
  - tabs/Tabs
    - handleChange
      - Tab の切り替え時
      - Tab の DataQueries を実行(actions.service.serviceGetQueries)
  - tables/IndexTable.tsx
    - handleLinkClick
      - リンク先表示のための DataQueries を実行する
  - view/View.tsx
    - handleSubmitOnSearchForm
      - 検索ボタンの Submit 時に、データ更新する

## Routing

- Service の特定は、URL パスによって行う
- Service の Root からは、searchParams によって Routing を行う
