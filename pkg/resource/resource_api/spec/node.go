package spec

type Node struct {
	Name     string
	State    string
	Warnings int
	Errors   int
	Labels   string
}

type GetNodes struct {
	Cluster string `validate:"required"`
}

type GetNodesData struct {
	Nodes []Node
}
