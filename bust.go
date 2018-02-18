package main

import (
	"os"

	"github.com/Jacobious52/bust/cmd/apply"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	app := kingpin.New("bust", "bust renders files with go templating rules")
	applyCommand.Configure(app)
	kingpin.MustParse(app.Parse(os.Args[1:]))
}
