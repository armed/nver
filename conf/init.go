package conf

import (
	"log"
	"os"
	"os/user"
	"runtime"
)

const (
	PATHVAR      = "NVER_PATH"
	DEFAULT_DIR  = "/.nver"
	BIN_DIR      = "/bin"
	VERSIONS_DIR = "/versions"

	WORK_DIR_INIT_ERR = "Could not initialize working directory: %v"
)

var (
	workDirPath = os.Getenv(PATHVAR)
)

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
			log.Fatalf(WORK_DIR_INIT_ERR, path)
		}
	} else if err != nil {
		log.Fatalf(WORK_DIR_INIT_ERR, path)
	}
}

func Init() {
	if runtime.GOOS != "darwin" && runtime.GOOS != "linux" {
		log.Fatalf("Sorry, nver currently supports only Mac OS X and Linux")
	}

	if workDirPath == "" {
		u, err := user.Current()
		if err != nil {
			log.Fatalf("Could not get current user information")
		}
		workDirPath = u.HomeDir + DEFAULT_DIR
	}

	checkDirs(workDirPath, VersionsPath())
}

func WorkPath() string {
	return workDirPath
}

func VersionsPath() string {
	return workDirPath + VERSIONS_DIR
}

func BinPath() string {
	return workDirPath + BIN_DIR
}
