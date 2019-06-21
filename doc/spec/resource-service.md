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
- ResourceMonitor
  - Node 群の監視をする Agent
  - 各 Node は一定間隔ごとに Api に自身のステータスを更新している
  - Monitor は、一定時間更新のない Node があった場合
    - その Node の Status API をたたき、その Node のステータスを更新する
  - その Node のステータス API をたたいて
- ResourceRegionApi
  - Region の Api
- ResourceRegionController
  - Resource を Region 単位でバッチ処理するコントローラ
  - Resource を ResourceAgent にアサインする
- ResourceAgent
  - Region に所属し、アサインされた Resource を実体化し、状態を管理する
  - Node ごとに複数の Provider サポートできる
  - Provider を利用してノード自身を監視し、イベントがあれば Alert を MonitorController に通知する
    - またメトリクスをメトリクス DB に送信する

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
    - GSLB の設定管理に利用
      - GSLB を有効にしなくてもよい
    - RegionServie を GSLB のメンバとして設定する
  - RegionService
    - GlobalService に紐ずく
      - GSLB を利用しない場合は紐図けなくてもよい
    - Region に紐ずく
      - 作成時に指定する
    - VIP の管理に利用
      - VIP を有効にしなくてもよい
    - 各種 Policy を設定する
      - Cluster: AntiAffinity
      - PowerAZ: AntiAffinity
      - RackAZ: AntiAffinity
      - NetworkAZ: AntiAffinity
  - 仮想リソースを作成する際は、RegionService を指定して作成する
    - オプションで Cluster、PowerAZ、RackAZ、NetworkAZ を指定する
    - RegionService の Policy によってリソースを配置する

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
