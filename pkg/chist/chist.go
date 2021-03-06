package chist

import (
	"fmt"
	"sort"
	"github.com/pkg/errors"
	"github.com/saromanov/chist/pkg/parser"
	"os"
)

type Pair struct {
	Command string
	Freq int
}

type PairList []Pair

func (p PairList) Len() int { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Freq < p[j].Freq }
func (p PairList) Swap(i, j int){ p[i], p[j] = p[j], p[i] }

// Chist defines the main structure of the app
type Chist struct {
	path string
	limit int
}

// New creates app
func New(path string, limit int) (*Chist, error) {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return nil, errors.New("file is not exists")
		}

		return nil, fmt.Errorf("unable to check file: %v", err)
	}
	return &Chist{
		path: path,
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
	pl := make(PairList, len(res))
	for k, v := range res {
		pl[i] = Pair{
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
