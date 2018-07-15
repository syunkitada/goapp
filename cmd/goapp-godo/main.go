package main

import (
	do "gopkg.in/godo.v2"
)

func tasks(p *do.Project) {
	p.Task("default", do.S{"goapp-db-migrate"}, nil)

	p.Task("goapp-db-migrate", nil, func(c *do.Context) {
		c.Bash("go run cmd/goapp-db-migrate/main.go")
	}).Src("pkg/model/**/*.go")

	p.Task("goapp-api", nil, func(c *do.Context) {
		c.Start("main.go", do.M{"$in": "cmd/goapp-api"})
	}).Src("pkg/**/*.go")

	p.Task("goapp-ctl", nil, func(c *do.Context) {
		c.Start("main.go", do.M{"$in": "cmd/goapp-ctl"})
	}).Src("pkg/**/*.go")
}

func main() {
	do.Godo(tasks)
}
