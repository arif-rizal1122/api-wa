package main

import (
	"api-wa/bootstrap"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	bootstrap.BootstrapApp()
}
