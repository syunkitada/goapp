package main

import (
	_ "github.com/go-sql-driver/mysql"

	"github.com/syunkitada/goapp/pkg/authproxy"
)

func main() {
	authproxy.Main()
}
