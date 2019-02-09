package main

import (
	_ "github.com/go-sql-driver/mysql"

	"github.com/syunkitada/goapp/pkg/ctl/ctl_admin"
)

func main() {
	ctl_admin.Main()
}
