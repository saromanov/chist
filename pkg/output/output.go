package output

import "github.com/saromanov/chist/pkg/models"

type Output interface {
	Show (pairs models.PairList) error
}