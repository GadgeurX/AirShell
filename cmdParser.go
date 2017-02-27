package main

import "regexp"
import "strings"

var delimiter string

// CmdParser is a cmd parser
type CmdParser struct {
	regexArg *regexp.Regexp
}

// Init initialize cmdParser
func (parser *CmdParser) Init() {
	parser.regexArg = regexp.MustCompile("'.+'|\".+\"|\\S+")
	delimiter = ";|<>"
}

// Parse parse a cmd line
func (parser *CmdParser) Parse(cmdLine string) []string {
	cmdArg := parser.regexArg.FindAllString(cmdLine, -1)
	parseSep(cmdArg)
	return cmdArg
}

func parseSep(execArgs []string) {
	cmds := []Cmd{}
	var cmd Cmd
	cmd.execArg = []string{}
	cmds = append(cmds, cmd)
	for _, arg := range execArgs {
		startCmd := 0
		for i, char := range arg {
			if strings.ContainsAny(string(char), delimiter) {
				cmds[len(cmds)-1].execArg = append(cmds[len(cmds)-1].execArg, arg[startCmd:i])
				cmds[len(cmds)-1].separator = byte(char)
				startCmd = i + 1
				cmds = append(cmds, Cmd{})
			}
		}
		cmds[len(cmds)-1].execArg = append(cmds[len(cmds)-1].execArg, arg[startCmd:len(arg)])
		cmds[len(cmds)-1].separator = 0
	}
}
