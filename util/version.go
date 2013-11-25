package util

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

var (
	validVersionSting    = regexp.MustCompile(`^([vV]{1}\d+\.\d+\.\d+(-[a-zA-Z]+-\d+)??){1}(\s+\d+\.\d+\.\d+(-[a-zA-Z]+-\d+)??)??\*??$`)
	validArgVersion      = regexp.MustCompile(`^[vV]??\d+\.\d+(\.\d+)??$`)
	validExactArgVersion = regexp.MustCompile(`^[vV]??\d+\.\d+\.\d+$`)

	ErrorVersionNotFound            = errors.New("Version not found")
	ErrorCantParseVersionArg        = errors.New("Can't parse version argument")
	ErrorFullVersionMustBeSpecified = errors.New("Full version must be specified (e.g. 0.10.21)")
	ErrorNoNodeJsVersionSpecified   = errors.New("No Node.js version specified")
)

type version string

type VersionList interface {
	Add(verStr string)
	FindExact(verStr string) (string, error)
	FindNewest(verStr string) (string, error)
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

func (v *versionList) FindExact(verStr string) (bestMatch string, err error) {
	checked, err := checkVersionArgument(verStr, validExactArgVersion)
	if err == nil {
		bestMatch, err = v.findByPrefix(checked)
	}
	return
}

func (v *versionList) FindNewest(verStr string) (bestMatch string, err error) {
	checked, err := checkVersionArgument(verStr, validArgVersion)
	if err == nil {
		bestMatch, err = v.findByPrefix(checked)
	}
	return
}

func (v *versionList) Count() int {
	return len(v.vers)
}

func (v *versionList) Vers() []string {
	return v.vers
}

func (v *versionList) findByPrefix(verStr string) (bestMatch string, err error) {
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
	if bestMatch == "" {
		err = ErrorVersionNotFound
	}
	return
}

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

	selfMinor, e1 := strconv.Atoi(self[1])
	thatMinor, e2 := strconv.Atoi(that[1])
	selfPatch, e1 := strconv.Atoi(self[2])
	thatPatch, e2 := strconv.Atoi(that[2])
	if e1 != nil || e2 != nil {
		panic(ErrorCantParseVersionArg)
	}

	return selfMinor*1000+selfPatch > thatMinor*1000+thatPatch
}

func checkVersionArgument(verStr string, reg *regexp.Regexp) (result string, err error) {
	result = strings.TrimSpace(strings.ToLower(verStr))
	if result == "" {
		err = ErrorNoNodeJsVersionSpecified
	}
	if !reg.MatchString(result) {
		if reg == validArgVersion {
			err = ErrorCantParseVersionArg
		} else {
			err = ErrorFullVersionMustBeSpecified
		}
	}
	return result, err
}

func extractVersionToken(verStr string) string {
	return strings.Split(strings.TrimSpace(verStr), " ")[0]
}
