package cmd

import (
	"fmt"
	"github.com/armed/nver/conf"
	"github.com/armed/nver/util"
	"github.com/codegangsta/cli"
	"log"
	"os"
)

func use(c *cli.Context) {
	validateArgsNum(c.Args(), 1)
	varArg := c.Args()[0]
	varArg = util.CheckVersionArgument(varArg)
	vList := installedVersions()
	success, bestMatch := vList.FindBest(varArg)
	if !success {
		log.Fatalf("Could not find any match for %v, is it installed?", varArg)
	}
	os.RemoveAll(conf.BinPath())
	os.Symlink(conf.VersionsPath()+"/"+bestMatch+"/bin", conf.BinPath())
	fmt.Printf("Now using %v\n", bestMatch)
}
