package main

import "github.com/saromanov/chist/pkg/parser"

var (
	zshHistoryPath = ""
)

func main(){
	if err := parser.Parse(zshHistoryPath); err != nil {
		panic(err)
	}
}