# Resource Service

- リソース管理サービス

## Overview

- ResourceApi
  - ResourceFlavor の管理
  - Resource の管理
  - Resource の作成リクエストなどをデータベースに保存する
- ResourceController
  - Resource をバッチ処理するコントローラ
  - Resource の作成リクエストを ResourceRegionApi にリクエストし伝搬させる
  - ResourceRegionApi から Resource 実態の状態を取得し、データベースに保存する
  - Node 群の監視
  - MonitorLoop
    - Node 群の監視をする
    - 各 Node は一定間隔ごとに Api に自身のステータスを更新している
- ResourceClusterApi
  - Cluster の Api
- ResourceClusterController
  - Resource を Cluster 単位でバッチ処理するコントローラ
  - Resource を ResourceAgent にアサインする
- ResourceClusterAgent
  - 各 Node 内で稼働する Agent
  - Region に所属し、アサインされた Resource を実体化し、状態を管理する
  - Node ごとに複数の Provider サポートできる
  - Provider を利用してノード自身を監視し、イベントがあれば Alert を MonitorController に通知する
    - またメトリクスをメトリクス DB に送信する
- ResourceClusterProxyAgent
  - 各 Node 内で稼働する Agent
  - マネジメントセグでリッスンする
- ResourceComputeClusterAgent
  - 各 ComputeVM 内で稼働する Agent
  - 各 ComputeVM は GRPC で ResourceClusterProxyAgent に、ログ、メトリクスを転送し、自身の Orchestration 情報を取得し、タスクを実行する
    - ResourceClusterProxyAgent はそのホスト上の VM のログ、メトリクスをまとめて ResourceClusterApi へ転送し、Orchestration 情報をまとめて取得する
    - ComputeVM は ResourceClusterProxyAgent と情報交換するためのローカルマネジメントセグを持つ
    - マネジメントセグは ComputeNode 内で閉じて管理され、1 ノードに対して 1 つのブリッジが 全ノードが同じセグを持っている
  - 初期化処理
    - マネジメントセグ情報や鍵情報は ConfigDrive に保存してマウントしておく
    - VM 起動時にネットワーク設定を行い、GRPC によって Orchestration 処理を開始する
  - Orchestration
    - ネットワークの設定
    - ブロックデバイスの設定
    - ユーザ、グループ、SSH 公開鍵などの設定
    - ローリングアップデート
    - L3DSR の設定
  - EtoE TLS
    - VM 単位でローカルに TLS プロキシサーバを立てて TLS 終端をする
      - 各 VM 上のサービスは 必ずこのプロキシサーバを経由してアクセスさせる
      - 共有の L7 プロキシを立てて裏は HTTP みたいなことはやらない
      - L7 プロキシを立てる場合は、L7 プロキシと TLS プロキシ間で別途 TLS 接続する
    - 証明書の自動更新をサポートする
    - トレーシングをサポートする

## Data Model

- ResourceFlavor
  - Name
  - Provider
  - Spec
- ComputeResource
  - Name
  - Provider
  - Labels
  - ScheduleRule
  - Provider
  - RequestMethod
  - Status
  - StatusReason
  - MaxRevisionHistory
  - ResourceRevisionID
- VolumeResource
  - Name
  - Provider
  - Labels
  - ScheduleRule
  - Provider
  - RequestMethod
  - Status
  - StatusReason
  - MaxRevisionHistory
  - ResourceRevisionID
- ImageResource
  - Name
  - Provider
  - Labels
  - ScheduleRule
  - Provider
  - RequestMethod
  - Status
  - StatusReason
  - MaxRevisionHistory
  - ResourceRevisionID
- LoadbalancerResource
  - Name
  - Provider
  - Labels
  - ScheduleRule
  - Provider
  - RequestMethod
  - Status
  - StatusReason
  - MaxRevisionHistory
  - ResourceRevisionID
- Provider
  - Name
  - Description
  - Kind
    - VirtualMachine, Pod, VirtualMachineDeployment, PodDeployment, Configmap, Secret, Loadbalancer, CI, CD
  - Driver
    - Libvirt, Docker, Vpp, Xdp, Etcd
- RegionAavailabilityZone
  - Name
- NetworkV4
- NetworkV4Port

## Region Data Model

- Resource
  - Name
  - Provider
  - Labels
  - ScheduleRule
  - Provider
  - RequestMethod
  - Status
  - StatusReason
  - MaxRevisionHistory
  - ResourceRevisionID
- ResourceRevision
  - ResourceFlavorID
  - Spec
- Node
  - Name
  - RegionAvaiabilityZone
  - NwAvaiabilityZone
  - NodeAvaiabilityZone
  - Labels
  - Providers
  - Schedulable
  - Status
  - StatusReason
  - EnableAutoCordon
  - AutoCordonInterval
  - EnableAutoDrain
  - AutoDrainInterval
  - EnableAlert
- ResourceAssignment
  - ResourceID
  - AgentID
  - Status
  - StatusReason
- Provider
  - Name
  - Description
  - Kind
    - VirtualMachine, Pod, VirtualMachineDeployment, PodDeployment, Configmap, Secret, Loadbalancer, CI, CD
  - Driver
    - Libvirt, Docker, Vpp, Xdp, Etcd

## ResourceApi Method

- CreateResourceFlavor
- UpdateResourceFlavor
- DeleteResourceFlavor
- ListResourceFlavor
- GetResourceFlavor
- CreateResource
- UpdateResource
- DeleteResource
- ListResource
- GetResource
- DeleteNode
- ListNode
- GetNode
- CordonNode
- UncordonNode
- DrainNode

## Api RPC Method

- UpdateNode

## Root と Datacenter と Cluster と NetworkAvailabilityZone と NodeAvailabilityZone

- Root
  - 最上位レイヤー
  - ここに API がありユーザのリクエストはすべてここに集約される
  - Controller により、各リクエストが非同期に処理される
- Datacenter・Floor・Rack
  - Datacenter は複数の Floor から構成され、Floor には複数の Rack が収容される
  - PhysicalResource の収容場所を管理するために利用される
  - 管理データとして存在するのみで、Datacenter ごとに特別な API などはない
- PhysicalResource
  - 物理リソースの管理単位
  - 物理リソースのメタ情報や、物理的な依存関係を管理する
  - 状態は持たず、静的な情報のみを管理する
- Cluster
  - Cluster は仮想リソースの所属単位で Datacenter に所属する
  - Cluster ごとに API があり、そのクラスタに所属するリソースを操作するために利用される
  - Cluster ごとの Controller により、各リクエストが非同期に処理される
- NetworkAvailabilityZone
  - Cluster のネットワーク冗長を考慮し、L3 管理レイヤ(コアルータ、アグリゲートルータなど)ごとに分割する
- NodeAvailabilityZone
  - Cluster のラック冗長、電源冗長を考慮し、Node の管理レイヤごとに分割する
- Node
  - Cluster に属し、PhysicalResource を一対一で管理するためのリソース
  - Status、State といった動的に変化する状態を持つ
  - 仮想リソースが実体化される場合、この Node に紐好いて管理される
- 仮想リソース
  - Compute, Network, Image などのリソース

## 処理フローの概要

- 全サービス共通の監視
  - Node は一定間隔ごとに Api をたたいて自身のステータスを更新する
  - GrpcApi
    - Status
      - 自身のステータスを返す
- ResourceApi
  - ResourceFlavor の管理
  - Resource の管理
  - Resource の作成リクエストなどを Master データベースに保存する
  - GrpcApi
    - ReassignRole
- ResourceController
  - MonitorLoop
    - Api から Node の一覧を取得する
      - 失敗する場合は、Alert を発生させるため Hook する
    - 自身の Role を決定する
      - Node 一覧から Master のステータスを見て、Active なら自身の Role は Node 一覧で取得したものとなる
      - Master のステータスが Down の場合、Master の再決定を行う
        - ResourceApi.ReassignRole を実行し、自身の Role を取得する
    - 自身が Master なら、Active 数が一定数以上なら処理をスキップする(Slave にまかせる)
      - 一定数以上がつねに全ノードの監視をバッチ的に行う
    - Node の一覧から長期間更新のない Node は、Down にする
    - Node の一覧から中期間更新のない Node は、StatusApi をたたいて Node の状態を更新する
  - MainLoop
    - 自身の Role が Master なら処理を行う
    - Master データベースから Resource 作成リクエスト作成し ResourceClusterApi に伝搬させる
    - ResourceClusterApi から Resource 実態の状態を取得し、Master データベースを更新する
- ResourceClusterApi
  - Cluster の Api
- ResourceClusterController
  - Resource を Cluster 単位でバッチ処理するコントローラ
  - Resource を ResourceAgent にアサインする
- ResourceAgent
  - Cluster に所属し、アサインされた Resource を実体化し、状態を管理する
  - Node ごとに複数の Provider サポートできる
  - Provider を利用してノード自身を監視し、イベントがあれば Alert を MonitorController に通知する
    - またメトリクスをメトリクス DB に送信する

## リソースの割り当て、課金について

- Overcommit はしない
- リソースは作成された時点でリザーブされ、リザーブされた時間によって課金額が決定する
- vcpu リソース
  - CPU 時間の制限をつけた vcpu、もしくは pinning された vpu が利用可能で、課金額が変わる
  - 機種によっても課金額を設定できるようにする
- disk, volume, network などの IO リソース
  - IO の制限をかけ、その種類によって課金額がかわる
  - 機種によっても課金額を設定できるようにする
- 物理マシンリソース
  - 物理マシンをそのまま提供するようなことはしないが、1 物理マシンを 1VM or 2VM セット(NUMA ごと)で占有して提供
    - ノイジーネイバの影響がないため、制限はつけない
  - 物理マシン自体に課金額を設定できるようにする

## ユーザサービスの冗長性について

- 障害範囲とリソース配置
  - 物理的やオペミスなどでダウンするのは PhysicalResource のみ
  - すべての仮想リソースは PhysicalResource に紐好くので、PhysicalResource から障害箇所が特定できる
    - ユーザが利用するのはすべて仮想リソース
    - ユーザが物理リソースを利用する場合も透過的に見せるだけで仮想リソースとして利用してもらう
  - PhysicalResource のダウンパターンとリソース配置
    - パターン 1: 災害などにより Datacenter ごと落ちる
      - Datacenter ごとに VIP を持ち、GSLB により拠点間冗長するようにリソース配置する
      - ユーザの裁量で各 Datacenter に所属する Cluster でリソース作成を行い、各 VIP を GSLB に紐図ける
    - パターン 2: Datacenter の電力供給元や UPS などのトラブルにより、その電源系統の PDU がすべて落ちる
      - PowerAvailabilityZone が分散するようにリソース配置する(AntiAffinityPowerAZ)
      - PowerAvailabilityZone が集中するようにリソース配置する(Affinity)
        - オプションで Power 名を指定できる
        - GSLB により冗長担保する
      - どちらかの Policy を必ず選択する
      - Policy によりクラスタ単位で、PowerAZ を考慮してリソース配置する
    - パターン 3: PDU などのトラブルにより、Rack ごと落ちる
      - RackAvailabilityZone が分散するようにリソース配置する(AntiAffinity)
      - RackAvailabilityZone が同じになるようにリソース配置する(Affinity)
        - オプションで Rack 名を指定できる
        - 同セグ通信の場合は、エッジスイッチでの折り返しとなり低レイテンシが期待できる
    - パターン 4: ネットワーク機器のトラブルによりその機器配下がすべてダウンする
      - ネットワーク機器側で基本的に冗長が取れてるものだが、そのトラブル時は基本的にわりとどうしようもない
      - どうしようもないケース
        - オペミスにより一部設定が消されたり、上書きされる
        - アクティブ・スタンバイで機器のフェイルオーバーに失敗し全断
        - アクティブ・スタンバイでフェイルオーバーしたがペアとなってる機器に設定が漏れており、一部断
        - 一部ポート不良で、通信が不安定となるがアクティブのままとなる
      - ダウンする範囲は、Datacenter ごと落ちる、複数の NetworkAZ ごと落ちる、単一 NetworkAZ が落ちる場合がある
        - Datacenter ごと落ちる場合は、パターン 1 に当てはまるのでここでは考慮しない
        - NetworkAZ が分散するようにリソースを配置する(AntiAffinity)
          - ポリシー名: AntiAffinityNetworkAZ
        - NetworkAZ が集中するようにリソースを配置する(Affinity)
          - オプションで NetworkAZ 名を指定できる
          - GSLB により冗長担保する
- サービスの冗長性管理
  - サービスの SLA に応じて、Datacenter、Power、Rack、Network の観点で冗長を考慮する必要がある
  - GlobalService
    - 拠点間冗長を提供するためのもの
    - 有効にしなくてもよい
    - 複数拠点での冗長手段
      - GSLB
        - DNS ベースの負荷分散
        - ヘルスチェックの結果により、レコードを操作することで拠点間冗長する
        - メンバの VIP が落ちた場合、ヘルスチェックにより検知してレコードから VIP を削除する
        - あくまで DNS なので TTL が経過してキャッシュから消えるまでは VIP にアクセスされる可能性がある
      - 単一エニーキャスト IP によるバランシング
        - 複数拠点の各 L4 ロードバランサから同一の VIP を広報し、ルータに L4 ロードバランサへの経路を複数持たせる
        - 特定 IP への経路が複数ある場合はコストの低いほうが優先され、 同一コストの場合は ECMP によりバランシングされる
        - L4 ロードバランサは、TCP セッションは持たずにステートレスであるとする
          - L4 ロードバランサがダウンしても TCP セッションは破棄されない
        - L4 ロードバランサの配下に L7 ロードバランサを配置し、ドメイン名やパスによってサービスを特定してロードバランスする
  - RegionService
    - GlobalService に紐ずく
      - 拠点間冗長をしない場合は紐図けなくてもよい
    - Region に紐ずく
      - 作成時に指定する
    - VIP の管理に利用
      - VIP を有効にしなくてもよい
    - VirtualResource を作成する際は、RegionService を作成することによって間接的に作成する
    - 各 VirtualResource ごとに SchedulePolicy を設定し、Cluster や Node のスケジューリングに利用する
      - Replicas
        - 作成するリソース数
      - ClusterFilters
        - クラスタ名によりリソースを作成するクラスタをフィルタリングする
      - ClusterLabelFilters
        - ラベルによりリソースを作成するクラスタをフィルタリングする
      - NodeFilters
        - ノード名によりリソースを作成するクラスタをフィルタリングする
      - NodeLabelFilters
        - ラベルによりリソースを作成するノードをフィルタリングする
      - NodeLabelSoftUntiAffinities
        - 特定ラベルが設定してあるノードにできるだけ分散するようにスケジューリングする
      - NodeLabelSoftAffinities
        - 特定ラベルが設定してあるノードにできるだけ集中するようにスケジューリングする
      - NodeLabelHardUntiAffinities
        - 特定ラベルが設定してあるノードに必ず分散するようにスケジューリングする
        - ノードの空きがない場合は Error となる
      - NodeLabelHardAffinities
        - 特定ラベルが設定してあるノードに必ず集中するようにスケジューリングする
        - ノードの空きがない場合は Error となる
    - Cluster のスケジューリング
      - ResourceController がスケジューリングを担当
      - Cluster は Region に紐図いており、Region により Cluster はフィルタリングされる
      - ClusterFilters, ClusterLabelFilters によりクラスタはフィルタリングされる
      - Cluster に設定された Weight によってソート(数値の大きいほうが優先)される
      - Weight が同じ Cluster が複数ある場合は、Cluster 内のリソース空き容量によりソートされる
      - Cluster をまたいだスケジューリングはしない
        - PowerAvailabilityZone などが被る可能性がある
    - Cluster 内でのスケジューリング
      - ResourceClusterController がスケジューリングを担当
    - ユーザが意識する AZ
      - 論理 AZ を 3 つ以上用意し、デフォルトで AZ 分散する
      - リソース間の通信レイテンシ意識する場合は、AZ を指定しての起動もできるようにする

## 各種リソースの管理

- RegionService の作成
  - RegionService には、Compute、Loadbalancer を紐図けることができ、これらをまとめて作成し管理する
- Compute の作成
  - RegionService に紐づき、直接 Compute を作成することはできない
  - 新規 RegionService を作るか、既存 RegionService の Replica をインクリメントすることで作成される
- Compute の削除
  - 既存 RegionService の Replica をデクリメントする
    - 自動で末尾の Compute から削除される
  - Compute を直接指定して削除する
    - RegionService の Replica が自動でデクリメントされる
    - 0 になる場合は RegionService も削除される
- Loadbalancer の作成
  - RegionService に一つ紐づき、直接 Loadbalancer を作成することはできない
  - 新規 RegionService を作ることで作成される
- Loadbalancer の削除
  - 直接削除できない
  - RegionService を削除したときに削除される

## 各種リソースのステータスフロー

- root: RegionService

  - Initializing
    - controller: assign Computes to cluster
    - controller: assign Loadbalancer to cluster
    - controller: assign ip to Computes and create Computes
    - controller: assign ip to Loadbalancer and create Loadbalancer
    - controller: update RegionService status to Creating
  - Creating
    - controller: get Computes status and update RegionService status to Active
  - Active

- root: Compute

  - Initializing
    - controller: create Compute by cluster-api
    - controller: update Compute status to Creating:Scheduled
  - Creating:Scheduled
    - controller: get Compute status from cluster-api, and update compute status to Active
  - Active

- root: Loadbalancer

  - Initializing
    - controller: create Loadbalancer by cluster-api
    - controller: update Loadbalancer status to Creating:Scheduled
  - Creating:Scheduled
    - controller: get Loadbalancer status from cluster-api, and update compute status to Active
  - Active

- cluster: Compute

  - Initializing
    - cluster-controller: schedule Compute to nodes group by RegionService
    - cluster-controller: create Assignments
    - cluster-controller: update Compute status to Creating:Scheduled
  - Creating:Scheduled
    - cluster-agent: get assignments from cluster-api, and create Compute on node
    - cluster-agent: update assignments status by cluster-api
    - cluster-controller: update Compute status to Active
  - Active

- cluster: Loadbalancer

  - Initializing
    - cluster-controller: schedule Loadbalancer to nodes group by RegionService
    - cluster-controller: create Assignments
    - cluster-controller: update Loadbalancer status to Creating:Scheduled
  - Creating:Scheduled
    - cluster-agent: get assignments from cluster-api, and create Loadbalancer on node
    - cluster-agent: update assignments status by cluster-api
    - cluster-controller: update Loadbalancer status to Active
  - Active

## ダッシュボードにおけるラック図、ネットワーク図の見せ方

- ラック図
  - Datacenter ごとの Floor、Cluster を表示する
  - Floor を選択することで、その Floor 単位の物理ラック図を表示
    - Floor 情報から Rack、PhysicalResource を Join して表示
  - Cluster を選択することで、そのクラスタ単位の仮想ラック図を表示
    - Cluster 情報から、PhysicalResource、仮想リソースを Join して表示
- ネットワーク図
  - PhysicalResource、一部仮想リソースは複数のリンク情報を持つ
  - リンク情報は Port の組み合わせで表現される
    - 一つの Port は Network、IP、Mac の情報を持つ
  - あるノードを起点としてそのノードの所属するネットワーク図を表示する
    - Network に Port が紐ずいてるので、Network 指定で、Port とそれに紐ずくリソース一覧を取得し Network 図を作成する
  - 指定されたノードが L3 スイッチであった場合は、複数 Network が含まれるので複数の Network を表示する

## データセンタ内のネットワークイメージ

- 各インターネットプロバイダから回線を借りてゲートウェイルータ(gateway-router)を接続する
- ゲートウェイルータに各フロアへ接続するためのルータ(floor-spine-router)を接続する

```
 provider1                          --- root-1-floor-spine-router01
----------- root-1-gateway-router01 --- root-1-floor-spine-router02
                                    --- root-1-floor-spine-router03
                                    --- root-1-floor-spine-router04
                                    ...

 provider2                          --- root-2-floor-spine-router01
----------- root-2-gateway-router01 --- root-2-floor-spine-router02
                                    --- root-2-floor-spine-router03
                                    --- root-2-floor-spine-router04
                                    ...
```

- floor-spine-router からは、各フロアを束ねているルータ(floor-leaf-router)に接続する
- フロア名は、棟、階、フロア番号からなり、データセンタ内でユニーク
  - 1 棟目-1 階-1 フロアなら、floor-1-1-1

```
                            --- floor-1-1-1-floor-leaf-router01
                            --- floor-1-1-1-floor-leaf-router02
root-1-floor-spine-router01 --- floor-1-1-1-floor-leaf-router03
root-1-floor-spine-router02 --- floor-1-1-1-floor-leaf-router04
root-1-floor-spine-router03 --- floor-1-1-2-floor-leaf-router01
root-1-floor-spine-router04 --- floor-1-1-2-floor-leaf-router02
                            --- floor-1-1-2-floor-leaf-router03
                            --- floor-1-1-2-floor-leaf-router04
                            --- floor-1-2-1-floor-leaf-router01
                            --- floor-1-2-1-floor-leaf-router02
root-2-floor-spine-router01 --- floor-1-2-1-floor-leaf-router03
root-2-floor-spine-router02 --- floor-1-2-1-floor-leaf-router04
root-2-floor-spine-router03 --- floor-1-2-2-floor-leaf-router01
root-2-floor-spine-router04 --- floor-1-2-2-floor-leaf-router02
                            --- floor-1-2-2-floor-leaf-router03
                            --- floor-1-2-2-floor-leaf-router04
                            ...
```

- floor-leaf-router には、各ラックへ接続するためのルータ(rack-spine-router)を接続する

```
                                --- floor-1-1-1-rack-spine-router01
                                --- floor-1-1-1-rack-spine-router02
floor-1-1-1-floor-leaf-router01 --- floor-1-1-1-rack-spine-router03
floor-1-1-1-floor-leaf-router02 --- floor-1-1-1-rack-spine-router04
floor-1-1-1-floor-leaf-router03 --- floor-1-1-1-rack-spine-router05
floor-1-1-1-floor-leaf-router04 --- floor-1-1-1-rack-spine-router06
                                --- floor-1-1-1-rack-spine-router07
                                --- floor-1-1-1-rack-spine-router08
                                ...
```

- rack-spine-router からは、各ラックを束ねているルータ(rack-leaf-router)に接続する

```
                                --- floor-1-1-1-rack-1-1-rack-leaf-router01
                                --- floor-1-1-1-rack-1-1-rack-leaf-router02
floor-1-1-1-rack-spine-router01 --- floor-1-1-1-rack-1-2-rack-leaf-router01
floor-1-1-1-rack-spine-router02 --- floor-1-1-1-rack-1-2-rack-leaf-router02
floor-1-1-1-rack-spine-router03 --- floor-1-1-1-rack-1-3-rack-leaf-router01
floor-1-1-1-rack-spine-router04 --- floor-1-1-1-rack-1-3-rack-leaf-router02
floor-1-1-1-rack-spine-router05 --- floor-1-1-1-rack-2-1-rack-leaf-router01
floor-1-1-1-rack-spine-router06 --- floor-1-1-1-rack-2-1-rack-leaf-router01
floor-1-1-1-rack-spine-router07 --- floor-1-1-1-rack-2-2-rack-leaf-router02
floor-1-1-1-rack-spine-router08 --- floor-1-1-1-rack-2-2-rack-leaf-router02
                                --- floor-1-1-1-rack-2-3-rack-leaf-router03
                                --- floor-1-1-1-rack-2-3-rack-leaf-router03
                                ...
```

- 各ラックには、rack-leaf-router が 2 台あり、自身のラック配下のサーバと、ペアとなるラック配下のサーバにそれぞれ配線し冗長化する

```
floor-1-1-1-rack-1-1-rack-reaf-router01 --- rack-1-1-server1
floor-1-1-1-rack-1-1-rack-reaf-router02 --- rack-1-2-server2
floor-1-1-1-rack-1-2-rack-reaf-router01 --- rack-1-1-server1
floor-1-1-1-rack-1-2-rack-reaf-router02 --- rack-1-2-server2
                                        ...
```

- まとめて一本にすると

```
internet --- gateway-router --- floor-spine-router --- floor-leaf-router --- rack-spine-router --- rack-leaf-router --- server
```

## VM ネットワーク(L3 広報方式)

- VM の IP ごとに専用の netns を持つ
  - netns 管理の IP として、169.254.32.0/19(169.254.32.1 - 169.254.63.254: 8192 個) を利用する
    - 1 ノードで最大 8192 個の VM IP を管理することを想定する
    - ゲスト のゲートウェイには、169.254.1.1 を利用する
    - 169.254.1.2/32 を lo につけてメタデータサービスやレポジトリサービスを提供する
- 例: vmnetwork(172.16.0.0/24)
  - vm1(ip: 172.16.0.1/32) - netns (169.254.32.1/32) - host1(192.168.10.1/24) - router1
  - vm2(ip: 172.16.0.2/32) - netns (169.254.32.2/32) - host1(192.168.10.1/24) - router1
  - vm3(ip: 172.16.0.3/32) - netns (169.254.32.3/32) - host1(192.168.10.2/24) - router1
  - vm4(ip: 172.16.0.4/32) - netns (169.254.32.4/32) - host1(192.168.10.2/24) - router1
- VM と ResourceClusterProxyAgent との通信
  - VM は、169.254.1.2 に対してリクエストを行う
  - ResourceClusterComputeAgent は、src ip によってその VM の妥当性を確認する
    - L3 で VM の IP への経路が決まってるため、VM は IP の偽装ができない
- ACL
  - ACL は vm 専用の netns 内で iptables によって行う

## 物理機材の新陳代謝

- 物理機材には 5 年ほどの保守期限があるため、保守切れとなる前に利用を止め、新しいものに切り替える必要がある
- ネットワーク機器、gateway-router, floor-spine-router, floor-leaf-router ,rack-spine-router ,rack-leaf-router については、新しい機器を取り付け、古い機器の BGP を切ることで入れ替えることが可能
- サーバ機器については、Rack 単位で行い、その Rack 内のサーバ上リソースをすべて他の Rack にライブマイグレーションし、その Rack 単位で機材を総入れ替えする
  - Rack マイグレーション用に空の Rack を余分に確保しておく必要がある
    - 空 Rack は筐体交換用の倉庫として利用するとよい
    - 機材が壊れたらここから交換して、センドバックする
- リージョン単位の閉鎖およびマイグレーション
  - データセンタ自体の劣化のためリージョン単位で閉鎖することも考慮する必要がある
    - 閉鎖単位でリージョンを切る
      - 基本的には、リージョンに対してデータセンタが一つの想定だが、クラスタ単位で閉鎖する場合はリージョンに対してクラスタも一つにしておく
    - リソース利用ユーザにはリージョンだけを意識してもらう
  - tokyo2 というリージョンを開設し、tokyo1 というリージョンを閉鎖する場合は、ユーザには tokyo1 リージョンから tokyo2 リージョンへ移動してもらう必要がある
    - この場合、サービスが GlobalServcie に紐ずくという前提で以下の 2 通りの方法で移動することができる
      - tokyo2 でサービスを作り直して GlobalService に紐づけ、tokyo1 のサービスを GlobalService から外して削除する
      - 作り直しが難しい場合は、GlobalService から tokyo1 のサービスを外して停止し、tokyo2 へコールドマイグレーションしてサービスを再開して GlobalService に紐づける
    - GlobalService に紐図かないサービス(バッチなど)の場合は、停止を許容して tokyo1 のサービスを停止し tokyo2 へ移行する

## 監視と運用

- クラスタ内監視と通知
  - システム監視は ClusterAgent の Check によって行われ、Event を作成して ClusterApi に報告する
  - ClusterController が Event を集約して他のシステムに通知する
  - 運用者は通知されたものを適宜対応して、システムの健全性を担保する
- クラスタ間監視
  - クラスタ内監視は、クラスタ自体が止まると通知ができなくなるため、クラスタ間で互いを監視する必要がある
  - ClusterAgent は、他クラスタを監視して Event を作成することができる

## TimeSeriesData

- TimeSeriesData の種別と用途
  - 解析用
    - 大容量、短期(2 週間程度)保存
    - 用途
      - CPU、メモリ、ディスクなどのデータによる異常解析
      - リソースアサイン時の重みづけ
      - 短期的なリソースの利用率・稼働率の表示
  - 統計用
    - (短期的にみると)小容量、長期(数年)保存
    - 用途
      - 年ごと、月ごとのリソースの利用率・稼働率の表示
      - 月ごと、日ごとのエラーレートの遷移

## LogData

- LogData の種別と用途
  - 解析用
    - 大容量、短期(2 週間程度)保存
    - データは欠ける、重複する可能性がある
    - 用途
      - Error レート、Warning レートの表示
        - Error の一覧、Warning の一覧、その詳細表示
      - App、Host に絞ってのログ一覧表示
        - TraceId ごとにログを集約する(Error 数、Warning 数)を出し、展開できるようにする
      - TraceId によるトレース検索
  - 監査用
    - (短期的にみると)小容量、長期(数年)保存
    - データは消えてはならない、確実に記録されなければならない
    - リクエスト処理上のトランザクション上でログを記録する
    - MySQL などのデータベースを利用する
    - アプリケーション側で担保すべきなのでサポートしない
  - システムサービスログ
    - システムサービスに対してユーザが行ったアクションのログ
    - リソースの作成や削除など
  - システムイベント
    - システムが自動的に行ったログ
    - ライブマイグレーション、リバランス、ホストのメンテナンスなど

## EventData

- Event は、障害の発生や、障害の復旧を通知するためのもの
- EventData の種別
  - Success
    - 正常
  - Critical
    - 即時に対応が必要
  - Warning
    - 即時に対応は必要はない、システム側で自動復旧する可能性もある
- Event の発生と保存
  - Event は Node の Check によって発生する(ReportNode)
  - Event は TimeSeriesData, LogData, プラグインから、計算して発生される
    - プラグインは、Nagios や Sensu などと同様に、スクリプトの実行結果から Event を発生させる
  - Check の Occurences 設定により Event の発生は抑制される
    - ReissueDuration も Event に埋め込む
  - ReportNode によって報告された Event は、ClusterApi により EventDatabase に保存される
- Event の抑制と配信
  - ClusterController は、Event を EventDatabase から取得して、EventRule に従って処理する
  - 一度処理された Event は IssuedEvents に追加され、一定時間(ReissueDuration)までは 同一 Event は無視される
    - 一定時間を経過しても Event を検知した場合は再度 EventRule に従って処理される

## EventRuleData

- EventRuleData の種別
  - Filter
    - Event 保存前にフィルタリングするためのルール
    - ルールに引っかかった Event は、データベースに保存せずに破棄する
  - Silenced
    - Event の Action をさせないためのルール
    - ルールに引っかかった Event は、何もせずに Silenced フラグを立てて IssuedEvent に追加して Event 処理を終了する
  - Aggregate
    - Event を集約するためのルール
    - Event は集約されてから Action によって処理される
  - Action
    - 特定の Action を実行するためのルール
    - Action の定義は設定ファイルによって管理される
    - Action の種類
      - CreateAlertData
      - DeleteAlertData
      - ApiHook

## AlertData

- AlertData の種別
  - Fatal
    - 即時に対応が必要
    - ユーザーサービスに影響が出ている
    - ユーザも閲覧できるもの
  - Critical
    - 即時に対応が必要
  - Warning
    - 対応は必要だが、即時の対応は必要ない
- Alert の発行
  - Event のハンドラにより Alert を発行することができる
    - 同一 Event による Alert の重複発行は抑制する(時間のみ更新)
  - Alert は運用者が必ずクローズする必要がある
  - 特定の Log Event は、Alert ベースでの対応が必要となる
    - Log の Event の場合、一定時間で Log が追えなくなるため同一 Log が発生しないと Success となってしまう

## MaintenanceData

- MaintenanceData の発行
  - 基本的には運用者がメンテナンス時に発行し、閉じるもの
    - システムが自動でメンテナンスをする場合もある
- メンテナンスは影響がでないと思われるものでも、想定外のことがあれば影響が出ることもある
  - 影響が出た場合に、それを把握できるよう本番環境で誰が何をしているかを把握するためのもの
- オンメンテができない場合、ユーザに通知してメンテ時に影響が出ないよう対応してもらう必要がある
