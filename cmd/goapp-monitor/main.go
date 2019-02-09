package main

import (
	_ "github.com/go-sql-driver/mysql"

	"github.com/syunkitada/goapp/pkg/monitor"
)

func main() {
	monitor.Main()
}
