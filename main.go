package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	initWorkDir()

	app := cli.NewApp()
	app.Name = "nver"
	app.Usage = "Manage your Node.js versions"
	app.Commands = []cli.Command{
		{
			Name:      "ls-remote",
			ShortName: "lsr",
			Usage:     "List all available Note.js versions",
			Action:    lsRemote,
		},
		{
			Name:   "ls",
			Usage:  "List all installed Node.js versions",
			Action: lsLocal,
		},
	}
	app.Run(os.Args)
}

func lsRemote(c *cli.Context) {

	versions := getVersions()

	for _, v := range versions {
		fmt.Println(*v)
	}
}

func lsLocal(c *cli.Context) {

}
