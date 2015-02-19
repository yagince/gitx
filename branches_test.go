package main

import (
	"testing"
)

func TestBranchesUp(t *testing.T) {
	branches := Branches{values: []string{"hoge", "foo"}}

	branches.up()
	if branches.selected != 0 {
		t.Error("selectedが0の場合はupしても0")
	}

	branches.selected = 10
	branches.up()
	if branches.selected != 9 {
		t.Error("selectedが1以上の場合は-1される")
	}
}

func TestBranchesDown(t *testing.T) {
	branches := Branches{values: []string{"hoge", "foo"}}

	branches.down()
	if branches.selected != 1 {
		t.Error("selectedが0の場合はdownすると1")
	}

	branches.selected = 1
	branches.down()
	if branches.selected != 1 {
		t.Error("selectedが末尾の場合はdownしても変わらない")
	}
}
