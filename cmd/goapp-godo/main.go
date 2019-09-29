package main

import (
	do "gopkg.in/godo.v2"
)

func tasks(p *do.Project) {
	p.Task("goapp-authproxy-api", nil, func(c *do.Context) {
		c.Start("main.go api", do.M{"$in": "cmd/goapp-authproxy"})
	}).Src("pkg/**/*.go")

	p.Task("goapp-dashboard", nil, func(c *do.Context) {
		c.Start("main.go", do.M{"$in": "cmd/goapp-dashboard"})
	}).Src("pkg/**/*.go")

	p.Task("goapp-monitor-api", nil, func(c *do.Context) {
		c.Start("main.go api", do.M{"$in": "cmd/goapp-monitor"})
	}).Src("pkg/**/*.go")

	p.Task("goapp-monitor-alert-manager", nil, func(c *do.Context) {
		c.Start("main.go alert-manager", do.M{"$in": "cmd/goapp-monitor"})
	}).Src("pkg/**/*.go")

	p.Task("goapp-monitor-agent", nil, func(c *do.Context) {
		c.Start("main.go agent", do.M{"$in": "cmd/goapp-monitor"})
	}).Src("pkg/**/*.go")

	p.Task("goapp-resource-api", nil, func(c *do.Context) {
		c.Start("main.go api", do.M{"$in": "cmd/goapp-resource"})
	}).Src("pkg/**/*.go")

	p.Task("goapp-resource-api2", nil, func(c *do.Context) {
		c.Start("main.go api --config-file config2.toml", do.M{"$in": "cmd/goapp-resource"})
	}).Src("pkg/**/*.go")

	p.Task("goapp-resource-api3", nil, func(c *do.Context) {
		c.Start("main.go api --config-file config3.toml", do.M{"$in": "cmd/goapp-resource"})
	}).Src("pkg/**/*.go")

	p.Task("goapp-resource-controller", nil, func(c *do.Context) {
		c.Start("main.go controller", do.M{"$in": "cmd/goapp-resource"})
	}).Src("pkg/**/*.go")

	p.Task("goapp-resource-controller2", nil, func(c *do.Context) {
		c.Start("main.go controller --config-file config2.toml", do.M{"$in": "cmd/goapp-resource"})
	}).Src("pkg/**/*.go")

	p.Task("goapp-resource-controller3", nil, func(c *do.Context) {
		c.Start("main.go controller --config-file config3.toml", do.M{"$in": "cmd/goapp-resource"})
	}).Src("pkg/**/*.go")

	p.Task("goapp-resource-cluster-api", nil, func(c *do.Context) {
		c.Start("main.go cluster-api", do.M{"$in": "cmd/goapp-resource"})
	}).Src("pkg/**/*.go")

	p.Task("goapp-resource-cluster-controller", nil, func(c *do.Context) {
		c.Start("main.go cluster-controller", do.M{"$in": "cmd/goapp-resource"})
	}).Src("pkg/**/*.go")

	p.Task("goapp-resource-cluster-agent", nil, func(c *do.Context) {
		c.Start("main.go cluster-agent", do.M{"$in": "cmd/goapp-resource"})
	}).Src("pkg/**/*.go")
}

func main() {
	do.Godo(tasks)
}
