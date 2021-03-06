package main

import "github.com/saromanov/chist/pkg/parser"

var (
	zshHistoryPath = "$GOPATH/.zsh_history"
)

func main(){
	if err := parser.Parse(zshHistoryPath); err != nil {
		panic(err)
	}
}