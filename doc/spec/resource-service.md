# Resource Service
* リソース管理サービス


## Overview
* ResourceApi
    * ResourceFlavorの管理
    * Resourceの管理
* ResourceController
    * ResourceをResourceAgentにアサインする
* ResourceAgent
    * Nodeごとに複数のProviderサポートし、アサインされたResourceを実体化し、状態を管理する
    * Providerを利用してノード自身を監視し、イベントがあればAlertをMonitorControllerに通知する
        * またメトリクスをメトリクスDBに送信する
* MonitorController
    * Alertを発行する
        * メール通知
        * フック


## Data Model
* ResourceFlavor
    * Name
    * Provider
    * Spec
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
