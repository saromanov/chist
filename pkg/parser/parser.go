package parser

import (
	"bufio"
	"github.com/pkg/errors"
	"os"
	"strings"
)

// Parse provides parsing of history file
// and returns line to the f
func Parse(path string, f func(string)) error {
	file, err := os.Open(path)
	if err != nil {
		return errors.Wrap(err, "unable to open file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ";")
		if len(line) <= 1 {
			continue
		}
		res := strings.Split(line[1], " ")
		if len(res) == 0 {
			continue
		}
		f(res[0])
	}

	if err := scanner.Err(); err != nil {
		return errors.Wrap(err, "unable to scan file")
	}
	return nil
}
