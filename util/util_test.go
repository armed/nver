package util

import (
	"testing"
)

var body = []byte(`#node npm
v0.10.16 1.3.8
v0.10.17 1.3.8
v0.10.18 1.3.8
v0.10.19 1.3.11
v0.10.20 1.3.11
v0.10.21 1.3.11
v0.10.22 1.3.14
v0.11.0 1.2.15
v0.11.1 1.2.18
v0.11.2 1.2.21
v0.11.3 1.2.25
v0.11.4 1.3.4
v0.11.5 1.3.6
v0.11.6 1.3.8
v0.11.7 1.3.8
v0.11.8 1.3.13`)

func Test_tokensFromVersionString(t *testing.T) {
	expect := []string{"v01", "v02"}
	str := "v01 v02"
	success, tokens := tokensFromVersionString(str)
	if !success {
		t.Error("It should success")
	}
	if tokens[0] != expect[0] && tokens[1] != expect[1] {
		t.Errorf("Expected %v, but got %v", expect, tokens)
	}

	str = "#node npm"
	success, _ = tokensFromVersionString(str)
	if success {
		t.Error("It shouldn't success")
	}

	str = ""
	success, _ = tokensFromVersionString(str)
	if success {
		t.Error("It shouldn't success")
	}
}

func Test_parseVersions(t *testing.T) {
	versions := parseVersions(body)
	if len(versions.Vers()) == 0 {
		t.Error("Versions should be populates")
	}
	if versions.Vers()[8].stable {
		t.Errorf("Version %v should be unstable", versions.Vers()[8].String())
	}
	if !versions.Vers()[0].stable {
		t.Errorf("Version %v should be stable", versions.Vers()[0].String())
	}
}

func Test_makeUrl(t *testing.T) {
	expect := "http://nodejs.org/dist/v0.10.21/node-v0.10.21-darwin-x64.tar.gz"
	version := NewVersion("v0.10.21", "2.3.1")
	url := makeUrl(version, "darwin", "amd64")
	if url != expect {
		t.Errorf("Expected: %v, but got: %v", expect, url)
	}

	expect = "http://nodejs.org/dist/v0.10.21/node-v0.10.21-darwin-x86.tar.gz"
	url = makeUrl(version, "darwin", "386")
	if url != expect {
		t.Errorf("Expected: %v, but got: %v", expect, url)
	}

	expect = "http://nodejs.org/dist/v0.10.21/node-v0.10.21-linux-x86.tar.gz"
	url = makeUrl(version, "linux", "386")
	if url != expect {
		t.Errorf("Expected: %v, but got: %v", expect, url)
	}

	expect = "http://nodejs.org/dist/v0.10.21/node-v0.10.21-linux-x64.tar.gz"
	url = makeUrl(version, "linux", "amd64")
	if url != expect {
		t.Errorf("Expected: %v, but got: %v", expect, url)
	}
}

func Test_FindBestMatch(t *testing.T) {
	versions := parseVersions(body)

	success, version := versions.FindBest("v0.10")
	if !success {
		t.Errorf("Expected to be v0.10.22, but got nil")
	} else if version.String() != "v0.10.22" {
		t.Errorf("Expected to be v0.10.22, but got %v", version.String())
	}

	success, version = versions.FindBest("0.11")
	if !success {
		t.Errorf("Expected to be v0.11.8, but got nil")
	} else if version.String() != "v0.11.8" {
		t.Errorf("Expected to be v0.11.8, but got %v", version.String())
	}
}
