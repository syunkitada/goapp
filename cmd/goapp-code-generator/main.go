package main

import (
	"github.com/syunkitada/goapp/pkg/base/code_generator"

	authproxy_spec "github.com/syunkitada/goapp/pkg/authproxy/spec"
	resource_spec "github.com/syunkitada/goapp/pkg/resource/spec"
)

func main() {
	code_generator.Generate(&authproxy_spec.Spec)
	code_generator.Generate(&resource_spec.Spec)
	code_generator.GenerateStatusCodes()
}
