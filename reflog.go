package gitx

import (
	"fmt"
	"strings"
)

type Reflog struct {
	Hash      string
	History   string
	Operation string
}

func NewReflog(log string) (*Reflog, error) {
	splitted := strings.SplitN(log, " ", 3)
	if len(splitted) < 3 {
		return nil, fmt.Errorf("reflog parse failed. [ %s ]", log)
	}

	return &Reflog{
		Hash:      splitted[0],
		History:   strings.TrimSuffix(splitted[1], ":"),
		Operation: splitted[2],
	}, nil
}

func (log Reflog) String() string {
	return strings.Join([]string{log.Hash, log.History, log.Operation}, " ")
}
