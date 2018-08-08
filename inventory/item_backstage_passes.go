package inventory

type BackstagePasses struct {
	item
}

func (b *BackstagePasses) Update() {
	current := b.sellIn.Value()

	var change int
	if current > 10 {
		change = 1
	}
	if current <= 10 && current > 5 {
		change = 2
	}
	if current <= 5 && current > 0 {
		change = 3
	}
	if b.sellIn.Passed() {
		change = -b.quality.Value()
	}

	b.quality.Update(change)
	b.sellIn.Decrement()
}
