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

var (
	archs = map[string]string{
		"amd64": "x64",
		"386":   "x86",
	}
)

// GetVersions downloads version list from http://nodejs.org/dist/npm-versions.txt and
// converts it to VersionList structure
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

	return NewVersionListFromSlice(strings.Split(string(body), "\n"))
}

// GetDownloadUrl creates url string for specified version
func GetDownloadUrl(ver string) string {
	return makeUrl(ver, runtime.GOOS, runtime.GOARCH)
}

func makeUrl(ver string, os, arch string) string {
	path := strings.Join([]string{DIST_BASE_URL, ver, "node"}, "/")
	path = strings.Join([]string{path, ver, os, archs[arch]}, "-")
	return path + ".tar.gz"
}
