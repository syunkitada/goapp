# Spec


## System Overview
* [AuthProxy Service](auth-proxy-service.md)
    * 認証・認可を行うプロキシ
    * クライアントからのすべてのリクエストはこのプロキシを経由し、各種サービスにアクセスする
* [Object Service](object-service.md)
    * オブジェクトを管理する
    * オブジェクトはVMイメージを含む
* [IPAM Service](ipam-service.md)
    * ネットワーク、IPアドレスを管理する
    * ネットワークの定義、IPアドレスの払い出しを行う
* [Resource Service](resource-service.md)
    * リソースの実態を管理する
    * リソースとは、仮想マシン、コンテナ、ロードバランサ、CI、CDを含む
    * リソースを扱う各ノードにはAgentが起動しており、リソース実態の管理と、ノード自体の監視を行う
