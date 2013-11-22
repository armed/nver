package cmd

import (
	"fmt"
	"github.com/armed/nver/conf"
	"github.com/armed/nver/util"
	"github.com/codegangsta/cli"
	"log"
	"os"
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
	os.RemoveAll(c.BinPath())
	os.RemoveAll(c.VersionsPath() + "/" + bestMatch)
	fmt.Printf("Removed %v, run install/use to start using another version\n", bestMatch)
}
