package main

import (
	"fmt"
	"os"

	"rsc.io/quote"

	"github.com/realjf/blockchain_demo/cli"
)

func main() {
	defer os.Exit(0)

	fmt.Println(quote.Hello())

	cli := cli.CommandLine{}
	cli.Run()
}
