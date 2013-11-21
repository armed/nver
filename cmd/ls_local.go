package cmd

import (
	"fmt"
	"github.com/armed/nver/conf"
	"github.com/armed/nver/util"
	"github.com/codegangsta/cli"
	"io/ioutil"
	"log"
)

func lsLocal(c *cli.Context) {
	vList := installedVersions()

	if vList.Count() > 0 {
		for _, v := range vList.Vers() {
			fmt.Println(v)
		}
	} else {
		fmt.Println("No installed Node.js versions")
	}
}

func installedVersions() (vList util.VersionList) {
	vers, err := ioutil.ReadDir(conf.VersionsPath())
	if err != nil {
		log.Fatal("Could not read versions directory")
	}
	vList = util.NewVersionList()

	for _, v := range vers {
		vList.Add(v.Name())
	}
	return
}
