package util

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

var (
	validVersionSting = regexp.MustCompile(`^([vV]{1}\d+\.\d+\.\d+(-[a-zA-Z]+-\d+)??){1}(\s+\d+\.\d+\.\d+(-[a-zA-Z]+-\d+)??)??\*??$`)
	validArgVersion   = regexp.MustCompile(`^[vV]??\d+\.\d+(\.\d+)??$`)
)

type version string

func (v version) greaterThan(other string) bool {
	if string(v) == "" {
		return false
	}
	if other == "" {
		return true
	}

	self := strings.Split(string(v), ".")
	that := strings.Split(other, ".")

	if len(that) < len(self) {
		return true
	} else if len(that) > len(self) {
		return false
	}

	selfMinor, err := strconv.Atoi(self[1])
	thatMinor, err2 := strconv.Atoi(that[1])
	selfPatch, err := strconv.Atoi(self[2])
	thatPatch, err2 := strconv.Atoi(that[2])
	if err != nil || err2 != nil {
		log.Fatalf("Could not compare versions %v and %v", v, other)
	}

	return selfMinor*1000+selfPatch > thatMinor*1000+thatPatch
}

type VersionList interface {
	Add(verStr string)
	FindBest(verStr string) (bool, string)
	Count() int
	Vers() []string
}

type versionList struct {
	vers []string
}

func NewVersionList() VersionList {
	return &versionList{make([]string, 0)}
}

func NewVersionListFromSlice(vers []string) VersionList {
	vl := NewVersionList()
	for _, vs := range vers {
		vl.Add(vs)
	}
	return vl
}

func (v *versionList) Add(verStr string) {
	if validVersionSting.MatchString(verStr) {
		ver := extractVersionToken(verStr)
		v.vers = append(v.vers, ver)
	}
}

func (v *versionList) FindBest(verStr string) (success bool, bestMatch string) {
	verStr = CheckVersionArgument(verStr)
	if !strings.HasPrefix(verStr, "v") {
		verStr = "v" + verStr
	}

	for _, ver := range v.vers {
		if strings.HasPrefix(ver, verStr) && version(ver).greaterThan(bestMatch) {
			bestMatch = ver
		} else if bestMatch != "" {
			break
		}
	}
	success = bestMatch != ""
	return
}

func (v *versionList) Count() int {
	return len(v.vers)
}

func (v *versionList) Vers() []string {
	return v.vers
}

func CheckVersionArgument(verStr string) string {
	verStr = strings.TrimSpace(strings.ToLower(verStr))
	if verStr == "" {
		log.Fatalf("No Node.js version specified")
	}
	if !validArgVersion.MatchString(verStr) {
		log.Fatalf("Can't parse version argument")
	}
	return verStr
}

func extractVersionToken(verStr string) string {
	return strings.Split(strings.TrimSpace(verStr), " ")[0]
}
