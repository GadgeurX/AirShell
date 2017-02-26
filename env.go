package main

import (
	"bytes"
	"os"
	"strings"
)

//Env contains all environement related stuff
type Env struct {
	data map[string]string
	path Path
}

// Init initialize env object
func (env *Env) Init() {
	env.data = make(map[string]string)
	for _, line := range os.Environ() {
		result := strings.Split(line, "=")
		env.data[result[0]] = result[1]
	}
	env.path.Init(env.data["PATH"])
}

// GetEnv return the env char**
func (env *Env) GetEnv() []string {
	var envArray []string
	for key, value := range env.data {
		var buffer bytes.Buffer
		buffer.WriteString(key)
		buffer.WriteString("=")
		buffer.WriteString(value)
		envArray = append(envArray, buffer.String())
	}
	return envArray
}
