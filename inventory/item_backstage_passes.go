package inventory

type BackstagePasses struct {
	item
}

func (b *BackstagePasses) Update() {
	days := b.sellIn.Days()
	change := 1
	if days <= 10 && days > 5 {
		change = 2
	}
	if days <= 5 && days > 0 {
		change = 3
	}
	if b.sellIn.Passed() {
		change = -b.quality.Value()
	}
	b.quality.Update(change)

	b.sellIn.Decrement()
}
