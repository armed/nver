package conf

import (
	"github.com/armed/nver/util"
	"io/ioutil"
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
	util.VersionList
}

type configuration struct {
	workPath string
	current  string
	util.VersionList
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
		instance = &configuration{workPath: wp}
		instance.VersionList = util.NewVersionList()

		if out, err := exec.Command(wp+"/current/bin/node", "-v").Output(); err == nil {
			instance.current = strings.TrimSpace(string(out))
			instance.VersionList.Add(instance.current + "*")
		}

		infos, err := ioutil.ReadDir(instance.workPath)
		if err != nil {
			log.Fatalln("Could not read versions directory")
		}
		for _, fi := range infos {
			instance.VersionList.Add(fi.Name())
		}
	})
	return instance
}

func (conf *configuration) WorkPath() string {
	return conf.workPath
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
