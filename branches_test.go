package main

import (
	"testing"
)

func TestBranchesUp(t *testing.T) {
	branches := Branches{values: []string{"hoge", "foo"}}

	branches.Up()
	if branches.selected != 0 {
		t.Error("selectedが0の場合はupしても0")
	}

	branches.selected = 10
	branches.Up()
	if branches.selected != 9 {
		t.Error("selectedが1以上の場合は-1される")
	}
}

func TestBranchesDown(t *testing.T) {
	branches := Branches{values: []string{"hoge", "foo"}}

	branches.Down()
	if branches.selected != 1 {
		t.Error("selectedが0の場合はdownすると1")
	}

	branches.selected = 1
	branches.Down()
	if branches.selected != 1 {
		t.Error("selectedが末尾の場合はdownしても変わらない")
	}
}

func TestBranchesSelectedBranchName(t *testing.T) {
	branches := Branches{values: []string{"hoge", "foo"}}

	if branches.SelectedBranch() != "hoge" {
		t.Error("default selected branch is index:0")
	}

	branches.Down()
	if branches.SelectedBranch() != "foo" {
		t.Error("downすると次のブランチが選択される")
	}

	branches.Up()
	if branches.SelectedBranch() != "hoge" {
		t.Error("upすると前のブランチが選択される")
	}
}

func TestBranchesCurrentBranchName(t *testing.T) {
	branches := Branches{values: []string{"hoge", "foo"}, current: 1}

	if branches.CurrentBranch() != "foo" {
		t.Error("current is index:1")
	}

	branches.Down()
	if branches.CurrentBranch() != "foo" {
		t.Error("downしても変わらない")
	}

	branches.Up()
	if branches.CurrentBranch() != "foo" {
		t.Error("upしても変わらない")
	}
}
