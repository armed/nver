package cmd

import (
	"fmt"
	"github.com/armed/nver/conf"
	"github.com/codegangsta/cli"
	"io/ioutil"
	"log"
	"strings"
)

func lsLocal(c *cli.Context) {
	vers, err := ioutil.ReadDir(conf.VersionsPath())
	if err != nil {
		log.Fatal("Could not read versions directory")
	}

	if len(vers) > 0 {
		for _, v := range vers {
			if strings.HasPrefix(v.Name(), "v") {
				fmt.Println(v.Name())
			}
		}
	} else {
		fmt.Println("No installed Node.js versions")
	}
}
