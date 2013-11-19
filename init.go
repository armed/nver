package main

import (
	"os"
	"os/user"
)

const (
	PATHVAR      = "NVER_PATH"
	DEFAULT_DIR  = "/.nver"
	BIN_DIR      = "/bin"
	VERSIONS_DIR = "/versions"
)

var workDir = os.Getenv(PATHVAR)

func initWorkDir() {
	if workDir == "" {
		u, err := user.Current()

		if err != nil {
			fatalError("Could not get current user information")
		}

		workDir = u.HomeDir + DEFAULT_DIR
	}

	checkDirs(workDir, workDir+BIN_DIR, workDir+VERSIONS_DIR)
}
