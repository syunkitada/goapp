package spec

type NetworkV4 struct {
	Kind        string `validate:"required"`
	Name        string `validate:"required"`
	Description string
	Cluster     string `validate:"required"`
	Subnet      string `validate:"required"`
	StartIp     string `validate:"required"`
	EndIp       string `validate:"required"`
	Gateway     string `validate:"required"`
}
