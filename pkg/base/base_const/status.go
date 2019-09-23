package base_const

const (
	StatusEnabled   = "Enabled"
	StatusEnabling  = "Enabling"
	StatusDisabled  = "Disabled"
	StatusDisabling = "Disabling"

	StatusInitializing          = "Initializing"
	StatusActive                = "Active"
	StatusCreating              = "Creating"
	StatusCreatingInitialized   = "Creating:Initialized"
	StatusCreatingScheduled     = "Creating:Scheduled"
	StatusUpdating              = "Updating"
	StatusUpdatingScheduled     = "Updating:Scheduled"
	StatusUnknownActivating     = "Unknown:Activating"
	StatusRebalancingUnassigned = "Rebalancing:Unassigned"
	StatusDeleting              = "Deleting"
	StatusDeletingScheduled     = "Deleting:Scheduled"
	StatusDeleted               = "Deleted"
	StatusError                 = "Error"
)

const (
	StateUp      = "Up"
	StateDown    = "Down"
	StateUnknown = "Unknown"
)

const (
	RoleLeader = "Leader"
	RoleMember = "Member"
)
