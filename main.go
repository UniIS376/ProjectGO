package main

import (
	"blockchaincoin/cli"
	"blockchaincoin/db"
)

func main() {
	defer db.Close()
	cli.Start()
}
