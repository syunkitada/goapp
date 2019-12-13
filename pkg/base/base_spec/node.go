package base_spec

type NodeService struct {
	Name         string
	Kind         string
	Role         string
	Status       string
	StatusReason string
	State        string
	StateReason  string
	Labels       string
	Spec         interface{}
}

type GetNodeService struct {
	Name string `validate:"required"`
	Kind string `validate:"required"`
}

type GetNodeServiceData struct {
	NodeService NodeService
}

type GetNodeServices struct{}

type GetNodeServicesData struct {
	NodeServices []NodeService
}

type CreateNodeService struct {
	Spec string `validate:"required" flagKind:"file"`
}

type CreateNodeServiceData struct{}

type UpdateNodeService struct {
	NodeService
}

type UpdateNodeServiceData struct{}

type DeleteNodeService struct {
	Name string `validate:"required"`
	Kind string `validate:"required"`
}

type DeleteNodeServiceData struct{}

type DeleteNodeServices struct {
	Spec string `validate:"required" flagKind:"file"`
}

type DeleteNodeServicesData struct{}
