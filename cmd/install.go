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
	"math"
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"
	"strings"
)

func Install(c *cli.Context) {
	validateArgsNum(c.Args(), 1)
	install(util.CheckVersionArgument(c.Args()[0]), conf.Get())
}

func install(ver string, c conf.Configuration) {
	availableVerions := util.GetVersions()
	success, bestMatch := availableVerions.FindBest(ver)
	if !success {
		log.Fatalf("Could not find matched version %v", ver)
	}
	urlStr := util.GetDownloadUrl(bestMatch)

	u, err := url.Parse(urlStr)
	if err != nil {
		log.Fatalf("Could parse download URL %v", err)
	}

	verDirPath := c.WorkPath() + "/" + strings.Split(path.Base(u.Path), "-")[1]
	err = os.Mkdir(verDirPath, os.ModeDir|os.ModePerm)
	if os.IsExist(err) {
		log.Fatalf("Version %v is already installed", ver)
	}
	if err != nil {
		log.Fatalf("Could not create version directory %v", err)
	}

	fmt.Printf("Downloading %v...\n", bestMatch)
	tarBuff := tar.NewReader(unzip(download(urlStr)))
	untar(tarBuff, verDirPath)

	fmt.Printf("Version %v successifully installed\n", bestMatch)
	fmt.Printf("Run 'nver use %v' command to start using it\n", bestMatch)
}

func download(urlStr string) io.Reader {
	response, err := http.Get(urlStr)
	if err != nil {
		log.Fatalf("Error while downloading %v: %v", urlStr, err)
	}
	defer response.Body.Close()

	length, err := strconv.Atoi(response.Header["Content-Length"][0])
	if err != nil {
		length = 0
	}

	return writeAndLogChunks(length, response.Body)
}

func writeAndLogChunks(length int, src io.Reader) io.Reader {
	gzBuff := new(bytes.Buffer)
	buf := make([]byte, 32*1024)
	total := 0.0
	for {
		nr, er := src.Read(buf)
		if nr > 0 {
			gzBuff.Write(buf[0:nr])
			total += float64(nr)

			if length > 0 {
				fmt.Printf("%v%%\r", math.Floor(total/float64(length)*100))
			}
		}
		if er == io.EOF {
			break
		}
		if er != nil {
			log.Fatal(er)
		}
	}
	fmt.Println()
	return gzBuff
}

func unzip(gzBuff io.Reader) io.Reader {
	fmt.Println("Extracting...")
	gz, _ := gzip.NewReader(gzBuff)

	tarBuff := new(bytes.Buffer)
	io.Copy(tarBuff, gz)
	return tarBuff
}

func untar(t *tar.Reader, rootPath string) {
	for {
		hdr, err := t.Next()
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Fatalf("Error %v", err)
		}

		info := hdr.FileInfo()
		name := stripRoot(info.Name())

		if info.IsDir() {
			os.MkdirAll(rootPath+name, os.ModeDir|os.ModePerm)
			untar(t, rootPath)
		} else if hdr.Typeflag == tar.TypeSymlink {
			err := os.Symlink(hdr.Linkname, rootPath+name)
			if err != nil {
				log.Fatalf("Could not create symlink %v -> %v: %v", hdr.Linkname, name, err)
			}
		} else {
			file, err := os.Create(rootPath + name)
			if err != nil {
				log.Fatalf("Could not create file %v: %v", name, err)
			}

			file.Chmod(info.Mode())
			if _, err := io.Copy(file, t); err != nil {
				log.Fatalf("Could not write file %v: %v", name, err)
			}

			file.Close()
		}
	}
}

func stripRoot(filePath string) string {
	tokens := strings.Split(filePath, "/")
	tokens[0] = ""
	return strings.Join(tokens, "/")
}
