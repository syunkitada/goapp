package resource_cluster_model

const (
	KindResourceClusterController = "ResourceClusterController"
	KindResourceClusterApi        = "ResourceClusterApi"
	KindResourceClusterAgent      = "ResourceClusterAgent"
)

const (
	StatusEnabled   = "Enabled"
	StatusDisabled  = "Disabled"
	StatusDisabling = "Disabling"
	StateUp         = "Up"
	StateDown       = "Down"
)

const (
	RoleLeader = "Leader"
	RoleMember = "Member"
)
