package cmd

import (
	"fmt"
	"github.com/armed/nver/conf"
	"github.com/armed/nver/util"
	"github.com/codegangsta/cli"
	"log"
	"os"
)

func Use(c *cli.Context) {
	validateArgsNum(c.Args(), 1)
	use(util.CheckVersionArgument(c.Args()[0]), conf.Get())
}

func use(ver string, c conf.Configuration) {
	vList := lsLocal(c)
	success, bestMatch := vList.FindBest(ver)
	if !success {
		log.Fatalf("Could not find any match for %v, is it installed?", ver)
	}
	if found, current := c.CurrentVersion(); found {
		os.Rename(c.WorkPath()+"/current", c.WorkPath()+"/"+current)
	}
	os.Rename(c.WorkPath()+"/"+bestMatch, c.WorkPath()+"/current")
	fmt.Printf("Now using %v\n", bestMatch)
}
