package gitx

type Branches struct {
	values   []string
	current  int
	selected int
}

func (b *Branches) Up() int {
	if b.selected != 0 {
		b.selected -= 1
	}
	return b.selected
}

func (b *Branches) Down() int {
	if (b.selected + 1) < len(b.values) {
		b.selected += 1
	}
	return b.selected
}

func (b *Branches) SelectedBranch() string {
	return b.values[b.selected]
}

func (b *Branches) CurrentBranch() string {
	return b.values[b.current]
}
