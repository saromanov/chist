package table

import (
	"github.com/saromanov/tables"
	"github.com/saromanov/chist/pkg/models"
)

// Table defines of the table representation
type Table struct {

}

// New provides puppet initialization
func New()*Table {
	return &Table{}
}

// Show showing of the tablke
func (t *Table) Show(pairs models.PairList) error {
	tab := tables.New()
	tab.AddHeader("Command", "Count")
	for _, p := range pairs {
		tab.AddLine(p.Command, p.Freq)
	}
	tab.Build()
	return nil
}

