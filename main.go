package main

import (
	"github.com/petterwoooju/nomadcoin/cli"
	"github.com/petterwoooju/nomadcoin/db"
)

func main() {
	defer db.Close()
	cli.Start()
}
