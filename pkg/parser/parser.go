package parser

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"os"
)

// Parse provides parsing of history file
func Parse(path, data string) error {
	file, err := os.Open("/path/to/file.txt")
	if err != nil {
		return errors.Wrap(err, "unable to open file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return errors.Wrap(err, "unable to scan file")
	}
	return nil
}
