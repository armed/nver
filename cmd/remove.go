package cmd

import (
	"fmt"
	"github.com/armed/nver/conf"
	"github.com/codegangsta/cli"
	"log"
	"os"
)

func Remove(context *cli.Context) {
	validateArgsNum(context.Args(), 1)
	verArg := context.Args()[0]
	c := conf.Get()
	result, err := findVersionFolder(verArg, c)
	if err != nil {
		log.Fatal(err)
	}
	os.RemoveAll(c.WorkPath() + result.path)
	if result.isCurrent {
		fmt.Printf("Removed currently used version %v, run 'nve use' to switch to another version",
			result.version)
	} else {
		fmt.Printf("Removed version %v", result.version)
	}
}
