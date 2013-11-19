package main

import (
	"strconv"
	"strings"
)

type version struct {
	node   string
	npm    string
	stable bool
}

func newVersion(node, npm string) *version {
	num, err := strconv.Atoi(strings.Split(node, ".")[1])

	if err != nil {
		fatalError("Could not parse version %v", node)
	}

	stable := num%2 == 0
	return &version{node, npm, stable}
}

func (v version) String() string {
	return v.node + " " + v.npm
}
