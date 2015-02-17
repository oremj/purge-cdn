package main

import (
	"fmt"
	"os"

	"github.com/oremj/purge-cdn/Godeps/_workspace/src/github.com/codegangsta/cli"
	"github.com/oremj/purge-cdn/cdns/edgecast"
)

func purgeEdgecastCommand() cli.Command {
	cmd := cli.Command{
		Name:  "edgecast",
		Usage: "purges url from edgecast",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:   "token",
				Usage:  "Access token.",
				EnvVar: "EDGECAST_TOKEN",
			},
			cli.StringFlag{
				Name:   "account-id",
				Usage:  "Account id",
				EnvVar: "EDGECAST_ACCOUNT_ID",
			},
			cli.StringFlag{
				Name:  "url, u",
				Usage: "URL to purge (required)",
			},
		},
		Action: doPurgeEdgecast,
	}

	return cmd
}

func doPurgeEdgecast(c *cli.Context) {
	if c.String("url") == "" {
		cli.ShowSubcommandHelp(c)
		fmt.Println("--url is required")
		os.Exit(1)
	}

	edgecastAPI := edgecast.NewAPI(c.String("account-id"), c.String("token"))

	id, err := edgecastAPI.Purge(c.String("url"))
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	fmt.Print(id)
}
