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

