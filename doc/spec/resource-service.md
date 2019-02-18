# Resource Service
* リソース管理サービス


## Overview
* ResourceApi
    * ResourceFlavorの管理
    * Resourceの管理
    * Resourceの作成リクエストなどをデータベースに保存する
* ResourceController
    * Resourceをバッチ処理するコントローラ
    * Resourceの作成リクエストをResourceRegionApiにリクエストし伝搬させる
    * ResourceRegionApiからResource実態の状態を取得し、データベースに保存する
    * Node群の監視
    * MonitorLoop
        * Node群の監視をする
        * 各Nodeは一定間隔ごとにApiに自身のステータスを更新している
* ResourceMonitor
    * Node群の監視をするAgent
    * 各Nodeは一定間隔ごとにApiに自身のステータスを更新している
    * Monitorは、一定時間更新のないNodeがあった場合
        * そのNodeのStatus APIをたたき、そのNodeのステータスを更新する
    * そのNodeのステータスAPIをたたいて
* ResourceRegionApi
    * RegionのApi
* ResourceRegionController
    * ResourceをRegion単位でバッチ処理するコントローラ
    * ResourceをResourceAgentにアサインする
* ResourceAgent
    * Regionに所属し、アサインされたResourceを実体化し、状態を管理する
    * Nodeごとに複数のProviderサポートできる
    * Providerを利用してノード自身を監視し、イベントがあればAlertをMonitorControllerに通知する
        * またメトリクスをメトリクスDBに送信する


## Data Model
* ResourceFlavor
    * Name
    * Provider
    * Spec
* ComputeResource
    * Name
    * Provider
    * Labels
    * ScheduleRule
    * Provider
    * RequestMethod
    * Status
    * StatusReason
    * MaxRevisionHistory
    * ResourceRevisionID
* VolumeResource
    * Name
    * Provider
    * Labels
    * ScheduleRule
    * Provider
    * RequestMethod
    * Status
    * StatusReason
    * MaxRevisionHistory
    * ResourceRevisionID
* ImageResource
    * Name
    * Provider
    * Labels
    * ScheduleRule
    * Provider
    * RequestMethod
    * Status
    * StatusReason
    * MaxRevisionHistory
    * ResourceRevisionID
* LoadbalancerResource
    * Name
    * Provider
    * Labels
    * ScheduleRule
    * Provider
    * RequestMethod
    * Status
    * StatusReason
    * MaxRevisionHistory
    * ResourceRevisionID
* Provider
    * Name
    * Description
    * Kind
        * VirtualMachine, Pod, VirtualMachineDeployment, PodDeployment, Configmap, Secret, Loadbalancer, CI, CD
    * Driver
        * Libvirt, Docker, Vpp, Xdp, Etcd
* RegionAavailabilityZone
    * Name
* NetworkV4
* NetworkV4Port


## Region Data Model
* Resource
    * Name
    * Provider
    * Labels
    * ScheduleRule
    * Provider
    * RequestMethod
    * Status
    * StatusReason
    * MaxRevisionHistory
    * ResourceRevisionID
* ResourceRevision
    * ResourceFlavorID
    * Spec
* Node
    * Name
    * RegionAvaiabilityZone
    * NwAvaiabilityZone
    * NodeAvaiabilityZone
    * Labels
    * Providers
    * Schedulable
    * Status
    * StatusReason
    * EnableAutoCordon
    * AutoCordonInterval
    * EnableAutoDrain
    * AutoDrainInterval
    * EnableAlert
* ResourceAssignment
    * ResourceID
    * AgentID
    * Status
    * StatusReason
* Provider
    * Name
    * Description
    * Kind
        * VirtualMachine, Pod, VirtualMachineDeployment, PodDeployment, Configmap, Secret, Loadbalancer, CI, CD
    * Driver
        * Libvirt, Docker, Vpp, Xdp, Etcd


## ResourceApi Method
* CreateResourceFlavor
* UpdateResourceFlavor
* DeleteResourceFlavor
* ListResourceFlavor
* GetResourceFlavor
* CreateResource
* UpdateResource
* DeleteResource
* ListResource
* GetResource
* DeleteNode
* ListNode
* GetNode
* CordonNode
* UncordonNode
* DrainNode


## Api RPC Method
* UpdateNode


## RootとDatacenterとClusterとNetworkAvailabilityZoneとNodeAvailabilityZone
* Root
    * 最上位レイヤー
    * ここにAPIがありユーザのリクエストはすべてここに集約される
    * Controllerにより、各リクエストが非同期に処理される
* Datacenter・Floor・Rack
    * Datacenterは複数のFloorから構成され、Floorには複数のRackが収容される
    * PhysicalResourceの収容場所を管理するために利用される
    * 管理データとして存在するのみで、Datacenterごとに特別なAPIなどはない
* Cluster
    * Clusterは仮想リソースの所属単位でDatacenterに所属する
    * ClusterごとにAPIがあり、そのクラスタに所属するリソースを操作するために利用される
    * ClusterごとのControllerにより、各リクエストが非同期に処理される
* NetworkAvailabilityZone
    * Clusterのネットワーク冗長を考慮し、L3管理レイヤ(コアルータ、アグリゲートルータなど)ごとに分割する
* NodeAvailabilityZone
    * Clusterのラック冗長、電源冗長を考慮し、ノードの管理レイヤごとに分割する


## 処理フローの概要
* 全サービス共通の監視
    * Nodeは一定間隔ごとにApiをたたいて自身のステータスを更新する
    * GrpcApi
        * Status
            * 自身のステータスを返す
* ResourceApi
    * ResourceFlavorの管理
    * Resourceの管理
    * Resourceの作成リクエストなどをMasterデータベースに保存する
    * GrpcApi
        * ReassignRole
* ResourceController
    * MonitorLoop
        * ApiからNodeの一覧を取得する
            * 失敗する場合は、Alertを発生させるためHookする
        * 自身のRoleを決定する
            * Node一覧からMasterのステータスを見て、Activeなら自身のRoleはNode一覧で取得したものとなる
            * MasterのステータスがDownの場合、Masterの再決定を行う
                * ResourceApi.ReassignRoleを実行し、自身のRoleを取得する
        * 自身がMasterなら、Active数が一定数以上なら処理をスキップする(Slaveにまかせる)
            * 一定数以上がつねに全ノードの監視をバッチ的に行う
        * Nodeの一覧から長期間更新のないNodeは、Downにする
        * Nodeの一覧から中期間更新のないNodeは、StatusApiをたたいてNodeの状態を更新する
    * MainLoop
        * 自身のRoleがMasterなら処理を行う
        * MasterデータベースからResource作成リクエスト作成しResourceClusterApiに伝搬させる
        * ResourceClusterApiからResource実態の状態を取得し、Masterデータベースを更新する
* ResourceClusterApi
    * ClusterのApi
* ResourceClusterController
    * ResourceをCluster単位でバッチ処理するコントローラ
    * ResourceをResourceAgentにアサインする
* ResourceAgent
    * Clusterに所属し、アサインされたResourceを実体化し、状態を管理する
    * Nodeごとに複数のProviderサポートできる
    * Providerを利用してノード自身を監視し、イベントがあればAlertをMonitorControllerに通知する
        * またメトリクスをメトリクスDBに送信する


## リソースの割り当て、課金について
* Overcommitはしない
* リソースは作成された時点でリザーブされ、リザーブされた時間によって課金額が決定する
* vcpuリソース
    * CPU時間の制限をつけたvcpu、もしくはpinningされたvpuが利用可能で、課金額が変わる
    * 機種によっても課金額を設定できるようにする
* disk, volume, networkなどのIO リソース
    * IOの制限をかけ、その種類によって課金額がかわる
    * 機種によっても課金額を設定できるようにする
* 物理マシンリソース
    * 物理マシンをそのまま提供するようなことはしないが、1物理マシンを1VM or 2VMセット(NUMAごと)で占有して提供
        * ノイジーネイバの影響がないため、制限はつけない
    * 物理マシン自体に課金額を設定できるようにする


## ユーザサービスの冗長性について
* 障害範囲とリソース配置
    * 物理的やオペミスなどでダウンするのはPhysicalResourceのみ
    * すべての仮想リソースはPhysicalResourceに紐好くので、PhysicalResourceから障害箇所が特定できる
        * ユーザが利用するのはすべて仮想リソース
        * ユーザが物理リソースを利用する場合も透過的に見せるだけで仮想リソースとして利用してもらう
    * PhysicalResourceのダウンパターンとリソース配置
        * パターン1: 災害などによりDatacenterごと落ちる
            * DatacenterごとにVIPを持ち、GSLBにより拠点間冗長するようにリソース配置する
            * ユーザの裁量で各Datacenterに所属するClusterでリソース作成を行い、各VIPをGSLBに紐図ける
        * パターン2: Datacenterの電力供給元やUPSなどのトラブルにより、その電源系統のPDUがすべて落ちる
            * PowerAvailabilityZoneが分散するようにリソース配置する(AntiAffinityPowerAZ)
            * PowerAvailabilityZoneが集中するようにリソース配置する(Affinity)
                * オプションでPower名を指定できる
                * GSLBにより冗長担保する
            * どちらかのPolicyを必ず選択する
            * Policyによりクラスタ単位で、PowerAZを考慮してリソース配置する
        * パターン3: PDUなどのトラブルにより、Rackごと落ちる
            * RackAvailabilityZoneが分散するようにリソース配置する(AntiAffinity)
            * RackAvailabilityZoneが同じになるようにリソース配置する(Affinity)
                * オプションでRack名を指定できる
                * 同セグ通信の場合は、エッジスイッチでの折り返しとなり低レイテンシが期待できる
        * パターン4: ネットワーク機器のトラブルによりその機器配下がすべてダウンする
            * ネットワーク機器側で基本的に冗長が取れてるものだが、そのトラブル時は基本的にわりとどうしようもない
            * どうしようもないケース
                * オペミスにより一部設定が消されたり、上書きされる
                * アクティブ・スタンバイで機器のフェイルオーバーに失敗し全断
                * アクティブ・スタンバイでフェイルオーバーしたがペアとなってる機器に設定が漏れており、一部断
                * 一部ポート不良で、通信が不安定となるがアクティブのままとなる
            * ダウンする範囲は、Datacenterごと落ちる、複数のNetworkAZごと落ちる、単一NetworkAZが落ちる場合がある
                * Datacenterごと落ちる場合は、パターン1に当てはまるのでここでは考慮しない
                * NetworkAZが分散するようにリソースを配置する(AntiAffinity)
                    * ポリシー名: AntiAffinityNetworkAZ
                * NetworkAZが集中するようにリソースを配置する(Affinity)
                    * オプションでNetworkAZ名を指定できる
                    * GSLBにより冗長担保する
* サービスの冗長性管理
    * サービスのSLAに応じて、Datacenter、Power、Rack、Networkの観点で冗長を考慮する必要がある
    * GlobalService
        * GSLBの設定管理に利用
            * GSLBを有効にしなくてもよい
        * ClusterServieをGSLBのメンバとして設定する
    * RegionService
        * GlobalServiceに紐ずく
            * GSLBを利用しない場合は紐図けなくてもよい
        * VIPの管理に利用
            * VIPを有効にしなくてもよい
        * 各種Policyを設定する
            * Cluster: AntiAffinity
            * PowerAZ: AntiAffinity
            * RackAZ: AntiAffinity
            * NetworkAZ: AntiAffinity
    * 仮想リソースを作成する際は、RegionとRegionServiceを指定して作成する
        * オプションでCluster、PowerAZ、RackAZ、NetworkAZを指定する
        * RegionServiceのPolicyによってリソースを配置する


## ダッシュボードにおけるラック図、ネットワーク図の見せ方
* ラック図
    * DatacenterごとのFloor、Clusterを表示する
    * Floorを選択することで、そのFloor単位の物理ラック図を表示
        * Floor情報からRack、PhysicalResource をJoinして表示
    * Clusterを選択することで、そのクラスタ単位の仮想ラック図を表示
        * Cluster情報から、PhysicalResource、仮想リソースをJoinして表示
* ネットワーク図
    * PhysicalResource、一部仮想リソースは複数のリンク情報を持つ
    * リンク情報はPortの組み合わせで表現される
        * 一つのPortはNetwork、IP、Macの情報を持つ
    * あるノードを起点としてそのノードの所属するネットワーク図を表示する
        * NetworkにPortが紐ずいてるので、Network指定で、Portとそれに紐ずくリソース一覧を取得しNetwork図を作成する
    * 指定されたノードがL3スイッチであった場合は、複数Networkが含まれるので複数のNetworkを表示する
