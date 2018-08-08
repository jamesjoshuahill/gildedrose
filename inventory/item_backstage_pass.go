package inventory

type backstagePass struct {
	item
}

func newBackstagePass(name string, sellIn, quality int) *backstagePass {
	return &backstagePass{item: newItem(name, sellIn, quality)}
}

func (b *backstagePass) Update() {
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
