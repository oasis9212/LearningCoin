package main

import (
	"ralo/cli"
	"ralo/db"
)

func main() {
	defer db.Close()
	cli.Start()
}
