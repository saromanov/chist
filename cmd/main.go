package main

import (
	"flag"
	"github.com/saromanov/chist/pkg/chist"
)

var (
	zshHistoryPath = ""
)

type config struct {
	FilePath string 
}

func parse() config{
	filePath := flag.String("path", "zsh", "path to zsh or bash history file")
	flag.Parse()
	return config{
		FilePath: filePath,
	}
}
func main() {
	cfg := parse()
	ch, err := chist.New(cfg.FilePath,10)
	if err != nil {
		panic(err)
	}
	if err := ch.Do(); err != nil {
		panic(err)
	}
}
