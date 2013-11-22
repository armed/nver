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
	os.RemoveAll(c.BinPath())
	os.Symlink(c.VersionsPath()+"/"+bestMatch+"/bin", c.BinPath())
	fmt.Printf("Now using %v\n", bestMatch)
}
