package cmd

import (
	"fmt"
	"github.com/armed/nver/conf"
	"github.com/codegangsta/cli"
)

func LsLocal(c *cli.Context) {
	vList := conf.Get()

	if vList.Count() > 0 {
		for _, v := range vList.Vers() {
			fmt.Println(v)
		}
	} else {
		fmt.Println("No installed Node.js versions")
	}
}
