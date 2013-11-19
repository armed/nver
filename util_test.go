package main

import (
	"testing"
)

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
	body := []byte(`#node npm
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

	versions := parseVersions(body)

	if len(versions) == 0 {
		t.Error("Versions should be populates")
	}

	if versions[8].stable {
		t.Errorf("Version %v should be unstable", versions[8].node)
	}

	if !versions[0].stable {
		t.Errorf("Version %v should be stable", versions[0].node)
	}
}
