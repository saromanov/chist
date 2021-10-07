package main

import (
	"flag"
	"github.com/saromanov/chist/pkg/chist"
)

var (
	zshHistoryPath = ""
)

func parse() chist.Config{
	filePath := flag.String("path", "zsh", "path to zsh or bash history file")
	flag.Parse()
	return chist.Config{
		FilePath: *filePath,
	}
}
func main() {
	cfg := parse()
	ch, err := chist.New(cfg,10)
	if err != nil {
		panic(err)
	}
	if err := ch.Do(); err != nil {
		panic(err)
	}
}
