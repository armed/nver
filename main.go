package main

import (
	"github.com/armed/nver/cmd"
	"github.com/armed/nver/conf"
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	conf.Init()

	app := cli.NewApp()
	app.Name = "nver"
	app.Version = "0.0.1"
	app.Usage = "Manage your Node.js versions"
	app.Commands = cmd.Commands
	app.Run(os.Args)
}
