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


## ApiとRegionApi
* ルートデータベース


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
