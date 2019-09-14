package spec

type Cluster struct {
	Region       string
	Datacenter   string
	Name         string
	Kind         string
	Description  string
	DomainSuffix string
	Labels       string
	Weight       int
}

type GetCluster struct {
	Name string
}

type GetClusterData struct {
	Cluster Cluster
}

type GetClusters struct{}

type GetClustersData struct {
	Clusters []Cluster
}
