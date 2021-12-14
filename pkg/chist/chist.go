package chist

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/pkg/errors"
	"github.com/saromanov/chist/pkg/models"
	"github.com/saromanov/chist/pkg/output"
	"github.com/saromanov/chist/pkg/parser"
)

// Chist defines the main structure of the app
type Chist struct {
	cfg Config
	path string
	limit int
	output output.Output
}

// New creates chist app
func New(cfg Config, limit int, output output.Output) (*Chist, error) {
	if cfg.FilePath == "" {
		return nil, errors.New("filepath is empty")
	}
	if _, err := os.Stat(cfg.FilePath); err != nil {
		if os.IsNotExist(err) {
			return nil, errors.New("file is not exists")
		}

		return nil, fmt.Errorf("unable to check file: %v", err)
	}
	return &Chist{
		path: cfg.FilePath,
		limit: limit,
		output: output,
		cfg: cfg,
	}, nil
}

// Do defines the main logic
func (c *Chist) Do() error {
	res := make(map[string]int)
	parser.Parse(c.path, func(line string){
		res[line]++
	}, func(line string) bool {
		if c.cfg.Contains != "" {
			return strings.Contains(line, c.cfg.Contains)
		}
		return true
	})
	i := 0
	pl := make(models.PairList, len(res))
	for k, v := range res {
		pl[i] = models.Pair{
			Command: k,
			Freq: v,
		}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	if len(pl) < c.limit {
		c.limit = len(pl)
	}
	c.output.Show(pl[0:c.limit])
	return nil
}
