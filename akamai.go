package main

import (
	"fmt"
	"os"

	"github.com/oremj/purge-cdn/Godeps/_workspace/src/github.com/codegangsta/cli"
	"github.com/oremj/purge-cdn/cdns/akamai"
)

func purgeAkamaiCommand() cli.Command {
	cmd := cli.Command{
		Name:  "akamai",
		Usage: "purges url from Akamai",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:   "user",
				Usage:  "Akamai username",
				EnvVar: "AKAMAI_USER",
			},
			cli.StringFlag{
				Name:   "password",
				Usage:  "Akamai password",
				EnvVar: "AKAMAI_PASSWORD",
			},
			cli.StringFlag{
				Name:  "url",
				Usage: "URL to purge (required)",
			},
		},
		Action: doPurgeAkamai,
	}

	return cmd
}

func doPurgeAkamai(c *cli.Context) {
	if c.String("url") == "" {
		cli.ShowSubcommandHelp(c)
		fmt.Println("--url is required")
		os.Exit(1)
	}

	akamaiAPI := akamai.NewAPI(c.String("user"), c.String("password"))

	resp, err := akamaiAPI.Purge(c.String("url"))
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	fmt.Println(resp.PurgeID)
}
