package cmd

import (
	"fmt"
	"github.com/armed/nver/conf"
	"github.com/armed/nver/util"
	"github.com/codegangsta/cli"
	"log"
	"os"
	"strings"
)

func Remove(c *cli.Context) {
	validateArgsNum(c.Args(), 1)
	remove(util.CheckVersionArgument(c.Args()[0]), conf.Get())
}

func remove(ver string, c conf.Configuration) {
	vList := lsLocal(c)
	success, bestMatch := vList.FindBest(ver)
	if !success {
		log.Fatalf("Could not find any match for %v, is it installed?", ver)
	}
	if strings.HasSuffix(bestMatch, "*") {
		fmt.Printf("Removed %v, run install/use to start using another version\n", bestMatch)
		os.RemoveAll(c.WorkPath() + "/current")
	} else {
		fmt.Printf("Removed %v", bestMatch)
		os.RemoveAll(c.WorkPath() + "/" + bestMatch)
	}
}
