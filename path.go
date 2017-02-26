package main

import (
	"strings"
)

// Path represente contains all path
type Path struct {
	paths []string
}

// Init initialize all path
func (path *Path) Init(pathString string) {
	path.paths = strings.Split(pathString, ":")
}
