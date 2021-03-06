package parser

import (
	"bufio"
	"fmt"
	"strings"
	"github.com/pkg/errors"
	"os"
)

// Parse provides parsing of history file
func Parse(path string) error {
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
		fmt.Println(line[1])
	}

	if err := scanner.Err(); err != nil {
		return errors.Wrap(err, "unable to scan file")
	}
	return nil
}

func calc()