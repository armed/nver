package util

import (
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"strings"
)

const (
	NODE_VERSIONS  = "http://nodejs.org/dist/npm-versions.txt"
	UNEXPECTED_ERR = "Unexpected error: %v"
	LSR_RESP_ERR   = "Error getting list of versions, response status %v"
	LSR_ERR        = "Error getting list of versions: %v"
	DIST_BASE_URL  = "http://nodejs.org/dist"
)

var archs = map[string]string{
	"amd64": "x64",
	"386":   "x86",
}

func GetVersions() VersionList {
	resp, err := http.Get(NODE_VERSIONS)
	if err != nil {
		log.Fatalf(UNEXPECTED_ERR, err)
	}

	if resp.StatusCode != 200 {
		log.Fatalf(LSR_RESP_ERR, resp.Status)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf(LSR_ERR, err)
	}

	return parseVersions(body)
}

func GetDownloadUrl(ver *version) string {
	return makeUrl(ver, runtime.GOOS, runtime.GOARCH)
}

func VersionFromString(ver string) (bestMatch *version) {
	if ver == "" {
		log.Fatalf("No Node.js version specified")
	}
	success, bestMatch := GetVersions().FindBest(ver)
	if !success {
		log.Fatalf("Could not find matched Node.js version")
	}
	return
}

func parseVersions(body []byte) VersionList {
	versions := NewVersionList()

	versionStrings := bodyToVersionStrings(body)
	for _, vs := range versionStrings {
		success, tokens := tokensFromVersionString(vs)
		if success {
			versions.Add(tokens[0], tokens[1])
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

func makeUrl(ver *version, os, arch string) string {
	path := strings.Join([]string{DIST_BASE_URL, ver.String(), "node"}, "/")
	path = strings.Join([]string{path, ver.String(), os, archs[arch]}, "-")
	return path + ".tar.gz"
}
