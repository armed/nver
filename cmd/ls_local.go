package cmd

import (
	"fmt"
	"github.com/armed/nver/conf"
	"github.com/armed/nver/util"
	"github.com/codegangsta/cli"
	"io/ioutil"
	"log"
)

func LsLocal(c *cli.Context) {
	vList := lsLocal(conf.Get())

	if vList.Count() > 0 {
		for _, v := range vList.Vers() {
			fmt.Println(v)
		}
	} else {
		fmt.Println("No installed Node.js versions")
	}
}

func lsLocal(c conf.Configuration) (vList util.VersionList) {
	vers, err := ioutil.ReadDir(c.VersionsPath())
	if err != nil {
		log.Fatal("Could not read versions directory")
	}
	vList = util.NewVersionList()

	for _, v := range vers {
		vList.Add(v.Name())
	}
	return
}
