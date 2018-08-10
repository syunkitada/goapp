package main

import (
	do "gopkg.in/godo.v2"
)

func tasks(p *do.Project) {
	p.Task("default", do.S{"goapp-adminctl-db-migrate"}, nil)

	p.Task("goapp-adminctl-db-migrate", nil, func(c *do.Context) {
		c.Bash("go run cmd/goapp-adminctl/main.go --use-testdata db-migrate")
		c.Bash("make compile-pb")
	}).Src("pkg/model/**/*.go")

	p.Task("goapp-authproxy", nil, func(c *do.Context) {
		c.Start("main.go --use-testdata", do.M{"$in": "cmd/goapp-authproxy"})
	}).Src("pkg/**/*.go")

	p.Task("goapp-health", nil, func(c *do.Context) {
		c.Start("main.go --use-testdata", do.M{"$in": "cmd/goapp-health"})
	}).Src("pkg/**/*.go")
}

func main() {
	do.Godo(tasks)
}
