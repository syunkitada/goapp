# IPAM Service


## Overview
* IpamApi


## Data Model
* Network
    * Name
    * Description
    * Type
        * flat, vlan, local
    * Provider
        * linuxbridge
* Subnet
    * SubnetはNetworkに紐ずく
    * NetworkID
    * Subnetmask
    * Gateway
    * IpRange
    * Resolv
    * EnableDhcp
    * EnableDns
    * EnableIpv6
* Port
    * SubnetID
    * IpAddress
    * MacAddress


## Method
* CreateNetwork
* UpdateNetwork
* DeleteNetwork
* ListNetwork
* GetNetwork
* CreateSubnetwork
* UpdateSubnetwork
* DeleteSubnetwork
* ListSubnetwork
* GetSubnetwork
* CreatePort
* UpdatePort
* DeletePort
* ListPort
* GetPort
