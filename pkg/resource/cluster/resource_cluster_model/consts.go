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

	StatusCreating = "Creating"
	StatusUpdating = "Updating"
	StatusDeleting = "Deleting"
)

const (
	RoleLeader = "Leader"
	RoleMember = "Member"
)
