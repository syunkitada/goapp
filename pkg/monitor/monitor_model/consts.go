package monitor_model

const (
	KindMonitorAlertManager = "MonitorAlertManager"
	KindMonitorApi          = "MonitorApi"
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
