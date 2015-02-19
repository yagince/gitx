package gitx

import (
	"testing"
)

func TestBranchesUp(t *testing.T) {
	branches := Branches{Values: []string{"hoge", "foo"}}

	branches.Up()
	if branches.Selected != 0 {
		t.Error("Selectedが0の場合はupしても0")
	}

	branches.Selected = 10
	branches.Up()
	if branches.Selected != 9 {
		t.Error("Selectedが1以上の場合は-1される")
	}
}

func TestBranchesDown(t *testing.T) {
	branches := Branches{Values: []string{"hoge", "foo"}}

	branches.Down()
	if branches.Selected != 1 {
		t.Error("Selectedが0の場合はdownすると1")
	}

	branches.Selected = 1
	branches.Down()
	if branches.Selected != 1 {
		t.Error("Selectedが末尾の場合はdownしても変わらない")
	}
}

func TestBranchesSelectedBranchName(t *testing.T) {
	branches := Branches{Values: []string{"hoge", "foo"}}

	if branches.SelectedBranch() != "hoge" {
		t.Error("default Selected branch is index:0")
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
	branches := Branches{Values: []string{"hoge", "foo"}, Current: 1}

	if branches.CurrentBranch() != "foo" {
		t.Error("Current is index:1")
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
