package base_spec

type NodeService struct {
	Name         string
	Kind         string
	Role         string
	Status       string
	StatusReason string
	State        string
	StateReason  string
	Token        string
	Endpoints    []string
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

type GetNodeServices struct {
	Name string
	Kind string
}

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
