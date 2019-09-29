package base_spec

type Node struct {
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

type GetNode struct {
	Name string `validate:"required"`
	Kind string `validate:"required"`
}

type GetNodeData struct {
	Node Node
}

type GetNodes struct{}

type GetNodesData struct {
	Nodes []Node
}

type CreateNode struct {
	Spec string `validate:"required" flagKind:"file"`
}

type CreateNodeData struct{}

type UpdateNode struct {
	Node
}

type UpdateNodeData struct{}

type DeleteNode struct {
	Name string `validate:"required"`
	Kind string `validate:"required"`
}

type DeleteNodeData struct{}

type DeleteNodes struct {
	Spec string `validate:"required" flagKind:"file"`
}

type DeleteNodesData struct{}
