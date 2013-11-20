package cmd

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"fmt"
	"github.com/armed/nver/conf"
	"github.com/armed/nver/util"
	"github.com/codegangsta/cli"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"
)

func install(c *cli.Context) {
	ver := strings.ToLower(c.Args()[0])
	if ver == "" {
		log.Fatalf("No Node.js version specified")
	}

	success, bestMatch := util.GetVersions().FindBest(ver)
	if !success {
		log.Fatalf("Could not find matched Node.js version")
	}

	urlStr := util.GetDownloadUrl(bestMatch)

	u, err := url.Parse(urlStr)
	if err != nil {
		log.Fatalf("Could acquire download URL %v", err)
	}

	verDirPath := conf.VersionsPath() + "/" + strings.Split(path.Base(u.Path), "-")[1]
	err = os.Mkdir(verDirPath, os.ModeDir|os.ModePerm)
	if os.IsExist(err) {
		log.Fatalf("Version %v is already installed", ver)
	}
	if err != nil {
		log.Fatalf("Could not create version directory %v", err)
	}

	tarBuff := tar.NewReader(unzip(download(urlStr)))
	untar(tarBuff, verDirPath)

	fmt.Printf("Version %v successifully installed\n", ver)
	fmt.Printf("Run 'nver use %v' command to start using it\n", ver)
}

func download(urlStr string) io.Reader {
	fmt.Println("Downloading...")
	response, err := http.Get(urlStr)
	if err != nil {
		log.Fatalf("Error while downloading %v: %v", urlStr, err)
	}
	defer response.Body.Close()

	gzBuff := new(bytes.Buffer)
	io.Copy(gzBuff, response.Body)
	return gzBuff
}

func unzip(gzBuff io.Reader) io.Reader {
	fmt.Println("Extracting...")
	gz, _ := gzip.NewReader(gzBuff)

	tarBuff := new(bytes.Buffer)
	io.Copy(tarBuff, gz)
	return tarBuff
}

func untar(tar *tar.Reader, path string) {
	for {
		hdr, err := tar.Next()
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Fatalf("Error %v", err)
		}

		info := hdr.FileInfo()
		name := stripRoot(info.Name())

		if info.IsDir() {
			os.MkdirAll(path+name, os.ModeDir|os.ModePerm)
			untar(tar, path)
		} else {
			file, err := os.Create(path + name)
			if err != nil {
				log.Fatalf("Could not create file %v", name)
			}

			file.Chmod(info.Mode())

			if _, err := io.Copy(file, tar); err != nil {
				log.Fatalf("Could not write file %v: %v", name, err)
			}
		}
	}
}

func stripRoot(filePath string) string {
	tokens := strings.Split(filePath, "/")
	tokens[0] = ""
	return strings.Join(tokens, "/")
}
