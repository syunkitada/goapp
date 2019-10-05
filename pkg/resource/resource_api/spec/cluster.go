package spec

type Cluster struct {
	Region       string `validate:"required"`
	Datacenter   string `validate:"required"`
	Name         string `validate:"required"`
	Kind         string `validate:"required"`
	Description  string
	DomainSuffix string `validate:"required"`
	Labels       string
	Weight       int
}

type GetCluster struct {
	Name string `validate:"required"`
}

type GetClusterData struct {
	Cluster Cluster
}

type GetClusters struct{}

type GetClustersData struct {
	Clusters []Cluster
}

type CreateCluster struct {
	Spec string `validate:"required" flagKind:"file"`
}

type CreateClusterData struct{}

type UpdateCluster struct {
	Name      string
	Region    string
	Token     string
	Endpoints []string
}

type UpdateClusterData struct{}

type DeleteCluster struct {
	Name string `validate:"required"`
}

type DeleteClusterData struct{}

type DeleteClusters struct {
	Spec string `validate:"required" flagKind:"file"`
}

type DeleteClustersData struct{}
