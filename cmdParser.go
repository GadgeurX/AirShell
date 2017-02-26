package main

import "regexp"

// CmdParser is a cmd parser
type CmdParser struct {
	regexArg *regexp.Regexp
}

// Init initialize cmdParser
func (parser *CmdParser) Init() {
	parser.regexArg = regexp.MustCompile("'.+'|\".+\"|\\S+")
}

// Parse parse a cmd line
func (parser *CmdParser) Parse(cmdLine string) []string {
	cmdArg := parser.regexArg.FindAllString(cmdLine, -1)
	return cmdArg
}
