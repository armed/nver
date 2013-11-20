package util

import (
	"strconv"
	"strings"
)

type version struct {
	node   string
	npm    string
	stable bool
}

func (v *version) String() string {
	return v.node
}

type VersionList interface {
	Add(node, npm string)
	FindBest(verStr string) (bool, *version)
	Vers() []*version
}

type versionList struct {
	vers []*version
}

func (v *versionList) Add(node, npm string) {
	v.vers = append(v.vers, NewVersion(node, npm))
}

func (v *versionList) FindBest(verStr string) (success bool, bestMatch *version) {
	if !strings.HasPrefix(verStr, "v") {
		verStr = "v" + verStr
	}

	for _, ver := range v.vers {
		if strings.HasPrefix(ver.node, verStr) {
			bestMatch = ver
		} else if bestMatch != nil {
			break
		}
	}

	success = bestMatch != nil

	return
}

func (v *versionList) Vers() []*version {
	return v.vers
}

func NewVersionList() VersionList {
	return &versionList{make([]*version, 0, 20)}
}

func NewVersion(node, npm string) *version {
	mmp := strings.Split(node, ".")

	minor, err := strconv.Atoi(mmp[1])

	if err != nil {
		FatalError("Could not parse version %v", node)
	}

	stable := minor%2 == 0
	return &version{strings.ToLower(node), npm, stable}
}
