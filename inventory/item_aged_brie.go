package inventory

type AgedBrie struct {
	item
}

func (a *AgedBrie) Update() {
	var change int
	if a.sellIn.Passed() {
		change = 2
	} else {
		change = 1
	}

	a.quality.Update(change)
	a.sellIn.Decrement()
}
