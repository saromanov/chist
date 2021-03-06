package chist

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/saromanov/chist/pkg/parser"
	"os"
)

// Chist defines the main structure of the app
type Chist struct {
	path string
}

// New creates app
func New(path string) (*Chist, error) {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return nil, errors.New("file is not exists")
		}

		return nil, fmt.Errorf("unable to check file: %v", err)
	}
	return &Chist{
		path: path,
	}, nil
}

// Do defines the main logic
func (c *Chist) Do() error {
	parser.Parse(c.path, func(line string){
		fmt.Println(line)
	})
	return nil
}
