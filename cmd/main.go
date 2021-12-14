package main

import (
	"flag"

	"github.com/saromanov/chist/pkg/chist"
	"github.com/saromanov/chist/pkg/output/table"
)


func parse() chist.Config{
	filePath := flag.String("path", "", "path to zsh or bash history file")
	contains := flag.String("contains", "", "contains part on string")
	flag.Parse()
	path := *filePath
	return chist.Config{
		FilePath: path,
		Contains: *contains,
	}
}
func main() {
	cfg := parse()
	tab := table.New()
	ch, err := chist.New(cfg,10, tab)
	if err != nil {
		panic(err)
	}
	if err := ch.Do(); err != nil {
		panic(err)
	}
}
