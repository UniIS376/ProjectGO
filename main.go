package main

import (
	"blockchaincoin/blockchain"
	"blockchaincoin/cli"
)

func main() {
	blockchain.Blockchain()
	cli.Start()
}
