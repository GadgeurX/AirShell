package main

import (
	"fmt"
)

// Shell contains all var for the shell
type Shell struct {
	env    Env
	reader CmdReader
	parser CmdParser
	exec   Executhor
}

// Init initialize the shell
func (shell *Shell) Init() {
	InitBuiltin()
	shell.env.Init()
	shell.reader.Init()
	shell.parser.Init()
	shell.exec.Init(&shell.env)
}

// Run the shell
func (shell *Shell) Run() {
	for {
		cmdLine, err := shell.reader.ReadCmdLine()
		if err != nil {
			fmt.Println(err)
		}
		cmdArg := shell.parser.Parse(cmdLine)
		if len(cmdArg) > 0 {
			shell.exec.Exec(cmdArg)
		}
	}
}
