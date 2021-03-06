package main

import "github.com/saromanov/chist/pkg/chist"

var (
	zshHistoryPath = ""
)

func main() {
	ch, err := chist.New(zshHistoryPath)
	if err != nil {
		panic(err)
	}
	if err := ch.Do(); err != nil {
		panic(err)
	}
}
