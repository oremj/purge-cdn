package main

import (
	"os"

	"github.com/oremj/purge-cdn/Godeps/_workspace/src/github.com/codegangsta/cli"
)

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
