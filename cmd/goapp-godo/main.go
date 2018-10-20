package main

import (
	do "gopkg.in/godo.v2"
)

func tasks(p *do.Project) {
	p.Task("default", do.S{"goapp-adminctl-db-migrate"}, nil)

	p.Task("goapp-adminctl-db-migrate", nil, func(c *do.Context) {
		c.Bash("go run cmd/goapp-adminctl/main.go db-migrate")
		c.Bash("make compile-pb")
	}).Src("pkg/model/**/*.go")

	p.Task("goapp-authproxy", nil, func(c *do.Context) {
		c.Start("main.go", do.M{"$in": "cmd/goapp-authproxy"})
	}).Src("pkg/**/*.go")

	p.Task("goapp-dashboard", nil, func(c *do.Context) {
		c.Start("main.go", do.M{"$in": "cmd/goapp-dashboard"})
	}).Src("pkg/**/*.go")

	p.Task("goapp-health", nil, func(c *do.Context) {
		c.Start("main.go", do.M{"$in": "cmd/goapp-health"})
	}).Src("pkg/**/*.go")

	p.Task("goapp-resource-api", nil, func(c *do.Context) {
		c.Start("main.go api", do.M{"$in": "cmd/goapp-resource"})
	}).Src("pkg/**/*.go")

	p.Task("goapp-resource-controller", nil, func(c *do.Context) {
		c.Start("main.go controller", do.M{"$in": "cmd/goapp-resource"})
	}).Src("pkg/**/*.go")

	p.Task("goapp-resource-region-server", nil, func(c *do.Context) {
		c.Start("main.go region-server", do.M{"$in": "cmd/goapp-resource"})
	}).Src("pkg/**/*.go")
}

func main() {
	do.Godo(tasks)
}
