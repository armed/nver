package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const (
	NODE_VERSIONS     = "http://nodejs.org/dist/npm-versions.txt"
	WORK_DIR_INIT_ERR = "Could not initialize working directory: %v"
	UNEXPECTED_ERR    = "Unexpected error: %v"
	LSR_RESP_ERR      = "Error getting list of versions, response status %v"
	LSR_ERR           = "Error getting list of versions: %v"
)

func fatalError(format string, a ...interface{}) {
	panic(fmt.Errorf(format, a))
}

func checkDirs(paths ...string) {
	if len(paths) > 0 {
		for _, p := range paths {
			checkDir(p)
		}
	}
}

func checkDir(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err = os.Mkdir(path, os.ModeDir|os.ModePerm); err != nil {
			fatalError(WORK_DIR_INIT_ERR, path)
		}
	} else if err != nil {
		fatalError(WORK_DIR_INIT_ERR, path)
	}
}

func getVersions() []*version {
	resp, err := http.Get(NODE_VERSIONS)

	if err != nil {
		fatalError(UNEXPECTED_ERR, err)
	}

	if resp.StatusCode != 200 {
		fatalError(LSR_RESP_ERR, resp.Status)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fatalError(LSR_ERR, err)
	}

	return parseVersions(body)
}

func parseVersions(body []byte) []*version {
	versionStrings := bodyToVersionStrings(body)

	versions := make([]*version, 0, len(versionStrings))

	for _, vs := range versionStrings {
		success, tokens := tokensFromVersionString(vs)

		if success {
			versions = append(versions, newVersion(tokens[0], tokens[1]))
		}
	}

	return versions
}

func bodyToVersionStrings(body []byte) []string {
	return strings.Split(string(body), "\n")
}

func tokensFromVersionString(vs string) (success bool, tokens []string) {
	tokens = strings.Split(vs, " ")
	success = len(tokens) == 2 && string(tokens[0][0]) != "#"
	return
}
