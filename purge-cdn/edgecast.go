package main

import (
	"fmt"
	"os"

	"github.com/oremj/go-purge-cdn/Godeps/_workspace/src/github.com/codegangsta/cli"
	"github.com/oremj/go-purge-cdn/cdns/edgecast"
)

func doPurgeEdgecast(c *cli.Context) {
	if c.String("url") == "" {
		cli.ShowSubcommandHelp(c)
		fmt.Println("--url is required")
		os.Exit(1)
	}

	edgecastAPI := &edgecast.API{
		AccountId: c.String("account-id"),
		Token:     c.String("token"),
	}

	id, err := edgecastAPI.Purge(c.String("url"))
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	fmt.Print(id)
}
