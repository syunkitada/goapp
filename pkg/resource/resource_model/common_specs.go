package resource_model

type NameSpec struct {
	Name string `validate:"required"`
}

type NetworkPolicySpec struct {
	Version        int
	Interfaces     int
	AssignPolicy   string
	StaticNetworks []string
}

type PortSpec struct {
	Version int
	Subnet  string
	Gateway string
	Ip      string
	Mac     string
}

type SchedulePolicySpec struct {
	Replicas                    int `validate:"required"`
	ClusterFilters              []string
	ClusterLabelFilters         []string
	NodeFilters                 []string
	NodeLabelFilters            []string
	NodeLabelHardAffinities     []string
	NodeLabelHardAntiAffinities []string
	NodeLabelSoftAffinities     []string
	NodeLabelSoftAntiAffinities []string
}
