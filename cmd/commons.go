package cmd

import (
	"github.com/armed/nver/util"
	"log"
	"strings"
)

type searchResults struct {
	path      string
	version   string
	isCurrent bool
}

func validateArgsNum(args []string, num int) {
	if len(args) < num {
		log.Fatal("Not enough arguments")
	}
}

func findVersionFolder(verArg string, vList util.VersionList) (result *searchResults, err error) {
	bestMatch, err := vList.FindExact(verArg)
	if err == nil {
		result = &searchResults{}
		result.version = bestMatch
		if result.isCurrent = strings.HasSuffix(bestMatch, "*"); result.isCurrent {
			result.path = "/current"
		} else {
			result.path = "/" + bestMatch
		}
	}
	return
}
