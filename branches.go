package gitx

type Branches struct {
	Values   []string
	Current  int
	Selected int
}

func (b *Branches) Up() int {
	if b.Selected != 0 {
		b.Selected -= 1
	}
	return b.Selected
}

func (b *Branches) Down() int {
	if (b.Selected + 1) < len(b.Values) {
		b.Selected += 1
	}
	return b.Selected
}

func (b *Branches) SelectedBranch() string {
	return b.Values[b.Selected]
}

func (b *Branches) CurrentBranch() string {
	return b.Values[b.Current]
}
