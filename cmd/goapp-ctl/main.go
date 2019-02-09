package main

import (
	_ "github.com/go-sql-driver/mysql"

	"github.com/syunkitada/goapp/pkg/ctl/ctl_main"
)

func main() {
	ctl_main.Main()
}
