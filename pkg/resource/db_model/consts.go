package db_model

const (
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
