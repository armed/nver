package main

import (
	"github.com/armed/nver/cmd"
	"github.com/codegangsta/cli"
	"log"
	"os"
	"runtime"
)

func main() {
	if runtime.GOOS != "darwin" && runtime.GOOS != "linux" {
		log.Fatalf("Sorry, nver currently supports only Mac OS X and Linux")
	}

	app := cli.NewApp()
	app.Name = "nver"
	app.Version = "0.0.1"
	app.Usage = "Manage your Node.js versions"
	app.Commands = cmd.Commands
	app.Run(os.Args)
}
