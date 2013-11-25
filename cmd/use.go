package cmd

import (
	"fmt"
	"github.com/armed/nver/conf"
	"github.com/codegangsta/cli"
	"log"
	"os"
)

func Use(c *cli.Context) {
	validateArgsNum(c.Args(), 1)
	use(c.Args()[0], conf.Get())
}

func use(verArg string, c conf.Configuration) {
	result, err := findVersionFolder(verArg, c)
	if err != nil {
		log.Fatal(err)
	}
	if !result.isCurrent {
		if found, current := c.CurrentVersion(); found {
			os.Rename(c.WorkPath()+"/current", c.WorkPath()+"/"+current)
		}
		os.Rename(c.WorkPath()+result.path, c.WorkPath()+"/current")
	}
	fmt.Printf("Now using %v\n", result.version)
}
