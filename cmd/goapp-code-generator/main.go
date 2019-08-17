package main

import (
	"github.com/syunkitada/goapp/pkg/authproxy/spec"
	"github.com/syunkitada/goapp/pkg/base/code_generator"
)

func main() {
	code_generator.Generate(&spec.Spec)
}
