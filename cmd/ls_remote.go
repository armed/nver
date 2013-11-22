package cmd

import (
	"fmt"
	"github.com/armed/nver/util"
	"github.com/codegangsta/cli"
)

func LsRemote(c *cli.Context) {
	versions := util.GetVersions()
	for _, v := range versions.Vers() {
		fmt.Println(v)
	}
}
