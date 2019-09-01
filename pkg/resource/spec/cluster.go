package spec

type GetCluster struct {
	Name string
}

type GetClusterData struct {
	Region       string
	Datacenter   string
	Name         string
	Kind         string
	Description  string
	DomainSuffix string
	Labels       string
}

type GetClusters struct{}

type GetClustersData struct {
	Clusters []GetClusterData
}
