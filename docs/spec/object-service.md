# Object Service
* オブジェクト配信サービス


## Overview
* ObjectApi


## Data Model
* Object
    * オブジェクトのメタデータ
        * オブジェクト名
        * ファイルフォーマット
        * ファイルの格納先
            * ローカルFS
            * URL
* Container
    * オブジェクトは必ず一つ以上のContainerに属する
    * コンテナのメタデータ
        * コンテナ名
        * プロジェクト名


## Method
* CreateContainer
* UpdateContainer
* DeleteContainer
* ListContainer
* GetContainer
* CreateObject
* UpdateObject
* DeleteObject
* ListObject
* GetObject
