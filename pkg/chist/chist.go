package chist

import (
	"fmt"
	"sort"
	"github.com/pkg/errors"
	"github.com/saromanov/chist/pkg/parser"
	"github.com/saromanov/chist/pkg/models"
	"os"
)


// Chist defines the main structure of the app
type Chist struct {
	path string
	limit int
}

// New creates chist app
func New(cfg Config, limit int) (*Chist, error) {
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
	}, nil
}

// Do defines the main logic
func (c *Chist) Do() error {
	res := make(map[string]int)
	parser.Parse(c.path, func(line string){
		res[line]++
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
	fmt.Println(pl[0:c.limit])
	return nil
}
