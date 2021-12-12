package models

// Pair defines pair of command and count of records
type Pair struct {
	Command string
	Freq int
}

type PairList []Pair

func (p PairList) Len() int { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Freq < p[j].Freq }
func (p PairList) Swap(i, j int){ p[i], p[j] = p[j], p[i] }