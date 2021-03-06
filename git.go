package gitx

import (
	"os/exec"
	"strconv"
	"strings"
)

type Git struct {
	binary    string
	directory string
}

func NewGit(directory string) *Git {
	var binary string
	var err error
	if binary, err = exec.LookPath("git"); err != nil {
		panic(err)
	}
	return &Git{binary: binary, directory: directory}
}

func (g *Git) Branches() *Branches {
	lf := "\n"
	cmd := exec.Command(g.binary, "branch")

	var out []byte
	var err error
	if out, err = cmd.Output(); err != nil {
		panic(err)
	}

	branches := strings.Split(strings.TrimRight(string(out), lf), lf)
	var current int
	for i, b := range branches {
		if strings.IndexAny(b, "*") == 0 {
			current = i
		}
	}
	return &Branches{Values: branches, Current: current}
}

func (g *Git) CheckOut(revision string) ([]byte, error) {
	revision = strings.Trim(revision, " *")
	return exec.Command(g.binary, "checkout", revision).CombinedOutput()
}

func (g *Git) Reflog(logNum int) []*Reflog {
	lf := "\n"

	options := []string{
		"reflog",
	}

	if logNum != 0 {
		options = append(options, "-n", strconv.Itoa(logNum))
	}

	out, err := exec.Command(g.binary, options...).CombinedOutput()
	if err != nil {
		panic(err)
	}

	logs := strings.Split(strings.TrimRight(string(out), lf), lf)
	reflogs := make([]*Reflog, len(logs))
	for i, log := range logs {
		reflog, err := NewReflog(log)
		if err != nil {
			panic(err)
		}
		reflogs[i] = reflog
	}
	return reflogs
}
