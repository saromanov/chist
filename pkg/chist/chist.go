package chist

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
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
	return &Check{
		path: path,
	}, nil
}

// Do defines the main logic
func (c *Chist) Do() error {

}