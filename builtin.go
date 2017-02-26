package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

type builtinFunc func([]string)

// Builtin contains all Builtin func
var Builtin map[string]builtinFunc

//InitBuiltin initialize builtinFunc
func InitBuiltin() {
	Builtin = make(map[string]builtinFunc)
	Builtin["exit"] = exit
	Builtin["env"] = env
	Builtin["setenv"] = setenv
	Builtin["unsetenv"] = unsetenv
	Builtin["cd"] = cd
}

func exit(execArg []string) {
	if len(execArg) > 1 {
		i, err := strconv.Atoi(execArg[1])
		if err != nil {
			fmt.Println(err)
		} else {
			os.Exit(i)
		}
	} else {
		os.Exit(0)
	}
}

func env(execArg []string) {
	for _, value := range mShell.env.GetEnv() {
		fmt.Println(value)
	}
}

func setenv(execArg []string) {
	if len(execArg) == 1 {
		env(execArg)
	} else if len(execArg) == 2 {
		mShell.env.data[execArg[1]] = ""
	} else {
		mShell.env.data[execArg[1]] = execArg[2]
	}
}

func unsetenv(execArg []string) {
	if len(execArg) > 1 {
		delete(mShell.env.data, execArg[1])
	}
}

func cd(execArg []string) {
	if len(execArg) == 2 {
		if execArg[1][0] != '/' {
			path, err := os.Getwd()
			var buffer bytes.Buffer
			buffer.WriteString(path)
			buffer.WriteString("/")
			buffer.WriteString(execArg[1])
			err = os.Chdir(buffer.String())
			if err != nil {
				fmt.Println(err)
			}
		} else {
			err := os.Chdir(execArg[1])
			if err != nil {
				fmt.Println(err)
			}
		}
	} else {
		err := os.Chdir(mShell.env.data["HOME"])
		if err != nil {
			fmt.Println(err)
		}
	}
}
