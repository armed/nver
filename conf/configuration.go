package conf

import (
	"log"
	"os"
	"os/user"
	"sync"
)

const (
	PATHVAR           = "NVER_PATH"
	DEFAULT_DIR       = "/.nver"
	BIN_DIR           = "/bin"
	VERSIONS_DIR      = "/versions"
	WORK_DIR_INIT_ERR = "Could not initialize working directory: %v"
)

var once sync.Once
var instance *configuration

type Configuration interface {
	WorkPath() string
	VersionsPath() string
	BinPath() string
}

type configuration struct {
	workDirPath string
}

func Get() Configuration {
	once.Do(func() {
		workDirPath := os.Getenv(PATHVAR)

		if workDirPath == "" {
			u, err := user.Current()
			if err != nil {
				log.Fatalf("Could not get current user information")
			}
			workDirPath = u.HomeDir + DEFAULT_DIR
		}
		instance = &configuration{workDirPath}
		checkDirs(workDirPath, instance.VersionsPath())
	})
	return instance
}

func (conf *configuration) WorkPath() string {
	return conf.workDirPath
}

func (conf *configuration) VersionsPath() string {
	return conf.workDirPath + VERSIONS_DIR
}

func (conf *configuration) BinPath() string {
	return conf.workDirPath + BIN_DIR
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
			log.Fatalf(WORK_DIR_INIT_ERR, path)
		}
	} else if err != nil {
		log.Fatalf(WORK_DIR_INIT_ERR, path)
	}
}
