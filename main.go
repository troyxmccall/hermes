

package main

import (
	"github.com/troyxmccall/hermes/cmd"
)

var (
	version            = "No version provided"
	commit             = "No commit provided"
	buildTime          = "No build timestamp provided"
)
//	app.Version = "Version:   " + version + "\n   Commit:    " + commit + "\n   BuildTime: " + buildTime

func main() {
	cmd.Execute()
}
