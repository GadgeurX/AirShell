package main

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"syscall"
)

//#include  <unistd.h>
//#include  <sys/types.h>
//#include  <stdio.h>
//#include  <sys/types.h>
//#include  <sys/wait.h>
import "C"

// Executhor execute a cmd
type Executhor struct {
	env *Env
}

// Init initialize Executhor
func (exec *Executhor) Init(env *Env) {
	exec.env = env
}

// Exec exec a cmd
func (exec *Executhor) Exec(execArg []string) {
	if exec.execBuiltin(execArg) == nil {
		return
	}
	path, err := exec.getBinaryPath(execArg[0])
	if err == nil {
		pid := C.fork()
		if pid != 0 {
			var status C.int
			C.wait(&status)
		} else {
			syscall.Exec(path, execArg, exec.env.GetEnv())
		}
	} else {
		fmt.Println(err)
	}
}

func (exec *Executhor) getBinaryPath(bin string) (string, error) {
	if _, err := os.Stat(bin); err == nil {
		return bin, nil
	}
	for _, path := range exec.env.path.paths {
		var buffer bytes.Buffer
		buffer.WriteString(path)
		if path[len(path)-1] != '/' {
			buffer.WriteString("/")
		}
		buffer.WriteString(bin)
		if _, err := os.Stat(buffer.String()); err == nil {
			return buffer.String(), nil
		}
	}
	return "", errors.New("go_shell: File not found")
}

func (exec *Executhor) execBuiltin(execArg []string) error {
	function, err := Builtin[execArg[0]]
	if !err {
		return errors.New("go_shell: Not a builtin")
	}
	function(execArg)
	return nil
}
