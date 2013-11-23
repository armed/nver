package conf

import (
	"log"
	"os"
	"os/exec"
	"os/user"
	"strings"
	"sync"
)

const (
	PATHVAR           = "NVER_PATH"
	DEFAULT_DIR       = "/.nver"
	WORK_DIR_INIT_ERR = "Could not initialize working directory: %v"
)

var once sync.Once
var instance *configuration

type Configuration interface {
	WorkPath() string
	CurrentVersion() (bool, string)
}

type configuration struct {
	workDir string
	current string
}

func Get() Configuration {
	once.Do(func() {
		wp := os.Getenv(PATHVAR)
		if wp == "" {
			u, err := user.Current()
			if err != nil {
				log.Fatalf("Could not get current user information")
			}
			wp = u.HomeDir + DEFAULT_DIR
		}
		checkDir(wp)
		instance = &configuration{workDir: wp}

		if out, err := exec.Command(wp+"/current/bin/node", "-v").Output(); err == nil {
			instance.current = strings.TrimSpace(string(out))
		}
	})
	return instance
}

func (conf *configuration) WorkPath() string {
	return conf.workDir
}

func (conf *configuration) CurrentVersion() (bool, string) {
	return conf.current != "", conf.current
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
