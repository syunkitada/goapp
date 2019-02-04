package main

import (
	_ "github.com/go-sql-driver/mysql"

	"github.com/syunkitada/goapp/pkg/resource"
)

func main() {
	resource.Main()
}
