package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

// CmdReader read cmd lines
type CmdReader struct {
}

// Init initialize cmdReader
func (reader *CmdReader) Init() {

}

// ReadCmdLine read a cmd line
func (reader *CmdReader) ReadCmdLine() (string, error) {
	bufReader := bufio.NewReader(os.Stdin)
	fmt.Print("go_shell $> ")
	var cmdLine []byte
	var cmd []byte
	var isPrefix bool
	var err error

	isPrefix = true
	for isPrefix {
		cmdLine, isPrefix, err = bufReader.ReadLine()
		if err != nil {
			return "", errors.New("go_shell: [ReadCmdLine] error while readLine")
		}
		cmd = append(cmd[:], cmdLine[:]...)
	}
	return string(cmd), nil
}
