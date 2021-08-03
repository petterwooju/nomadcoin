package main

import (
	"github.com/petterwoooju/nomadcoin/blockchain"
	"github.com/petterwoooju/nomadcoin/cli"
	"github.com/petterwoooju/nomadcoin/db"
)

func main() {
	defer db.Close()
	blockchain.Blockchain().AddBlock("First")
	blockchain.Blockchain().AddBlock("Second")
	blockchain.Blockchain().AddBlock("Third")
	cli.Start()
}
