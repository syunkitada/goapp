package spec

type Node struct {
	Name    string
	Kind    string
	State   string
	Wanings int
	Errors  int
	Labels  string
}

type GetNodes struct {
	Cluster string
}

type GetNodesData struct {
	Node []Node
}
