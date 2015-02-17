package main

import (
	"os"

	"github.com/oremj/go-purge-cdn/Godeps/_workspace/src/github.com/codegangsta/cli"
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

func main() {
	app := cli.NewApp()

	app.Name = "purge-cdn"
	app.Version = "0.1"
	app.Author = ""
	app.Email = ""
	app.Commands = []cli.Command{
		purgeEdgecastCommand(),
	}
	app.Run(os.Args)
}
