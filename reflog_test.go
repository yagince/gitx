package gitx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewReflog(t *testing.T) {
	log := "a703652 HEAD@{0}: commit: vendoring"
	reflog, _ := NewReflog(log)
	expect := &Reflog{Hash: "a703652", History: "HEAD@{0}", Operation: "commit: vendoring"}

	assert.Equal(t, expect, reflog)
}

func TestNewReflog_InvalidLog(t *testing.T) {
	log := "a703652 HEAD@{0}"
	_, err := NewReflog(log)
	assert.NotNil(t, err)
}
