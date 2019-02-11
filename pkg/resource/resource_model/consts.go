package resource_model

const (
	KindResourceController = "ResourceController"
	KindResourceApi        = "ResourceApi"
)

const (
	StatusEnabled   = "Enabled"
	StatusEnabling  = "Enabling"
	StatusDisabled  = "Disabled"
	StatusDisabling = "Disabling"

	StatusActive              = "Active"
	StatusCreating            = "Creating"
	StatusCreatingInitialized = "Creating:Initialized"
	StatusCreatingScheduled   = "Creating:Scheduled"
	StatusUpdating            = "Updating"
	StatusUpdatingScheduled   = "Updating:Scheduled"
	StatusDeleting            = "Deleting"
	StatusDeletingScheduled   = "Deleting:Scheduled"
	StatusDeleted             = "Deleted"
)

const (
	StateUp   = "Up"
	StateDown = "Down"
)

const (
	RoleLeader = "Leader"
	RoleMember = "Member"
)

const (
	SpecNetworkV4    = "NetworkV4"
	SpecCompute      = "Compute"
	SpecContainer    = "Container"
	SpecImage        = "Image"
	SpecVolume       = "Volume"
	SpecLoadbalancer = "Loadbalancer"
)

const (
	SpecKindNetworkV4Local  = "Local"
	SpecKindComputeLibvirt  = "Libvirt"
	SpecKindContainerDocker = "Docker"
	SpecKindImageUrl        = "Url"
	SpecKindVolumeNfs       = "Nfs"
	SpecKindVolumeIscsi     = "Iscsi"
	SpecKindLoadbalancerVpp = "Vpp"
)
