package monitor_model

type IgnoreAlertSpec struct {
	Index  string `validate:"required"`
	Host   string `validate:"required"`
	Name   string `validate:"required"`
	Level  string `validate:"required"`
	User   string `validate:"required"`
	Reason string `validate:"required"`
	Until  int64  `validate:"required"`
}
