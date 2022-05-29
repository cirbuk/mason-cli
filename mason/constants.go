package main

import (
	"embed"
	"fmt"
	"strings"
)

//go:embed help
var f embed.FS

var (
	Usage string = strings.TrimSpace(usage)
	//go:embed usage.txt
	usage string
)

func getHelpMessage(cmd string) string {
	switch cmd {
	case "usage": // if its the 'help' command
		return string(Usage)
	default: // for all others
		helpContent, _ := f.ReadFile(fmt.Sprintf("help/%s.txt", cmd))
		return (string(helpContent))
	}
}
