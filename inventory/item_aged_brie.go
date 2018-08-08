package inventory

type AgedBrie struct {
	item
}

func (a *AgedBrie) Update() {
	change := 1
	if a.sellIn.Passed() {
		change = 2
	}
	a.quality.Update(change)

	a.sellIn.Decrement()
}
