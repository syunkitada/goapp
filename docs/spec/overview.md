# Overview

## user to authproxy
```
user ----(public domain)----> GDNS
     <---------vip1----------

                      ----- authproxy(region1)
user ------> LB(vip1) ----- authproxy
                      ----> authproxy
                         |     |
                         | readonlydb(region1, replication from masterdb)
                         |
                      masterdb(multi region)
                         |
                         | readonlydb(region2, replication from masterdb)
                         |     |
                      ----- authproxy(region2)
     ------- LB(vip2) ----- authproxy
                      ----- authproxy
```

## authproxy to service
```
authproxy --(latency 1ms)--- service-api(region1)
          --(latency 1ms)--> service-api
                                 |    |
                                 | readonlydb (region1, replication from masterdb)
                                 |
                             service masterdb(multi region)
                                 |
                                 | readonlydb (region2, replication from masterdb)
                                 |    |
          --(latency 5ms)--- service-api(region2)
          --(latency 3ms)--- service-api
```

## resouce service
```
resource-api(multi region)
   |
   --> masterdb <--
                   |
     resource-controller(multi region)
        |
        |    cluster1             <------------ resource-cluster-agent(network az1, node az1)
        |--- resource-cluster-api <------------ resource-cluster-agent(network az1, node az1)
        |--- resource-cluster-api <------------ resource-cluster-agent(network az1, node az2)
        |--> resource-cluster-api <------------ resource-cluster-agent(network az2, node az1)
        |        |                <------------ resource-cluster-agent(network az2, node az2)
        |        |
        |        |--> cluster-db <--|
        |        |                  |
        |        |              resource-cluster-controller
        |        |
        |        |--> cluster-metrics-db <--|
        |        |--> cluster-logs-db    <--|
        |        |--> cluster-alerts-db  <--|
        |                                   |
        |     --------------->  resource-cluster-alert-controller ----- publish alerts ----> someone
        |     |
     resource-alert-controller(multi region) ------------ publish alerts ------->
        |
        |
        |    cluster2
        |--- resource-cluster-api
        |--- resource-cluster-api
        |--- resource-cluster-api
        |
```
