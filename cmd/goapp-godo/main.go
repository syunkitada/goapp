package main

import (
	do "gopkg.in/godo.v2"
)

func tasks(p *do.Project) {
	// Authproxy Services
	p.Task("goapp-authproxy-api", nil, func(c *do.Context) {
		c.Start("main.go api", do.M{"$in": "cmd/goapp-authproxy"})
	}).Src("pkg/**/*.go")

	// Dashboard Services
	p.Task("goapp-dashboard", nil, func(c *do.Context) {
		c.Start("main.go", do.M{"$in": "cmd/goapp-dashboard"})
	}).Src("pkg/**/*.go")

	// Home Services
	p.Task("goapp-home-api", nil, func(c *do.Context) {
		c.Start("main.go api", do.M{"$in": "cmd/goapp-home"})
	}).Src("pkg/**/*.go")

	p.Task("goapp-home-controller", nil, func(c *do.Context) {
		c.Start("main.go controller", do.M{"$in": "cmd/goapp-home"})
	}).Src("pkg/**/*.go")

	// Resource Services
	p.Task("goapp-resource-api", nil, func(c *do.Context) {
		c.Start("main.go api", do.M{"$in": "cmd/goapp-resource"})
	}).Src("pkg/**/*.go")

	p.Task("goapp-resource-controller", nil, func(c *do.Context) {
		c.Start("main.go controller", do.M{"$in": "cmd/goapp-resource"})
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
