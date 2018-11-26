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
	SpecCompute      = "Compute"
	SpecContainer    = "Container"
	SpecImage        = "Image"
	SpecVolume       = "Volume"
	SpecLoadbalancer = "Loadbalancer"
)

const (
	SpecComputeLibvirt  = "Libvirt"
	SpecContainerDocker = "Docker"
	SpecImageUrl        = "Url"
	SpecVolumeNfs       = "Nfs"
	SpecVolumeIscsi     = "Iscsi"
	SpecLoadbalancerVpp = "Vpp"
)
