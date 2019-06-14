# UI

```
Data でJsonで返す

クライアントは型を意識しない
返ってきた値から型を自動判定するのみ

データの型、UIなどのSpecはすべてサービス側に集約する
modelにどのようなリクエストを想定するかを書いておく
modelからドキュメント生成
modelからAPIひな型生成

yamlでviewとデータモデルを定義
cliインターフェイスはデータモデルだけで生成可能
viewはwebui生成用の補完メタデータ情報

サービスにspecディレクトリを作成し、そこにyamlを補完する


View:
    Kind: panels
    Panels:
      - Name: datacenters
        Kind: table
        Query: GetDatacenters
        Data: datacenter
      - Name: resources
        Type: tabs
        Tabs:
            - Name
              Kind: table
              Query: GetPhysicalResources
              Data: PhysicalResource
              Actions:
                - Name: Create
                  Icon: Create
                  Query: Create
                  Kind: Form
              SelectActions:
                - Name: Delete
                  Icon: Delete
                  Query: Delete
                  Kind: Form
      - Name: resource


# すべてcamel型
# パブリックな情報はすべて大文字で始める
# プライベートな情報はすべて小文字で始める
# cli時はgetDetailはget-detailに変換される
# cli時はオプションをyaml, json, コマンドオプションで指定できるようにする
    * 引数は必ず一つ以内
    * 引数にyaml, jsonが含まれる場合はファイルから読み、オプションで上書きする

Datacenter:
    Queries: [Create, Update, Get, GetDetail, Delete]
    Fields:
        Name:
            Type: string
            Gorm: not null;size:50;unique_index;
            View:
                IsSearch: true
                Link: "Datacenters/:0/Resources/Resources"
                LinkSync: true
                LinkGetQueries: [GetPhysicalResources, "GetRacks"]
            Create: true
            Get: true
            Option:
                GetDetail: required
        Kind
            Type: string
            Create: true
            Get: true
            Option:
                Get: option
        Description
            Type: string
            Create: true
            Update: true
            GetDetail: true

```
