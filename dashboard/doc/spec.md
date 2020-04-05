# Spec


## Design

#### Containers (src/containers)
* すべての根幹をつかさどるComponent群
* Root.tsx
  * RootのComponent
* AuthRoute.tsx
  * 認証を管理するためのComponent
  * 認証済みかどうかをチェックし、認証済みでないなら、ログイン用のURLへリダイレクトする


#### Components (src/components)
* UI表示のためのComponents群
* 基本的にステート情報の管理は行わない(Component内で閉じるようなステート管理は行う)
* 更新されたステート情報に基づいて愚直にUIを生成する
* componentは、他のcomponentを取り込む場合がある
* frames/Dashboard.tsx
  * UIの一番外側のフレームを管理するためのComponent
  * renderIndexで生成したRootのComponentを一つ持ち、表示する


#### Apps (src/apps)
* Action、Reducerの定義、処理を行うためのアプリケーション群
* 基本的にActionはComponentによって発行され、処理され、Reducerによってステートが更新される


#### Lib (src/lib)
* ライブラリ
* 汎用的な関数などを管理する



## ステートの更新タイミング
* 方針
  * 画面を表示するためのステートは事前に更新する必要がある
  * 画面遷移時に、Action(actions.service.serviceGetQueries)を実行して、ステートを更新する
    * Lodingのステートに移行する
  * 非同期で、Queriesが実行され、再度ステートが更新されて、データがそろうことで画面の表示を完了する
* 基本的なステートの更新タイミングは以下
  * 初回ロード時(認証直後)
  * Dashboardのプロジェクト切り替え時
  * Dashboardのサービス切り替え時
  * Componentsの特定イベント時
* Componentsによる更新タイミングは以下
  * panels/ExpansionPanels
    * componentWillMount
      * URLからgetQueriesを実行する
      * actions.service.serviceGetQueries
    * handleChange
      * URLの変化がないので、うまくやる
      * actions.service.serviceGetQueries
  * panes/Panes
    * componentWillMount
      * URLからgetQueriesを実行する
      * actions.service.serviceGetQueries
  * tabs/Tabs
    * handleChange
      * Tabの切り替え時
      * TabのDataQueriesを実行(actions.service.serviceGetQueries)
  * tables/IndexTable.tsx
    * handleLinkClick
      * リンク先表示のためのDataQueriesを実行する
  * view/View.tsx
    * handleSubmitOnSearchForm
      * 検索ボタンのSubmit時に、データ更新する
