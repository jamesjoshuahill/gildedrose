package inventory

type AgedBrie struct {
	item
}

func (a *AgedBrie) Update() {
	var change int
	if a.sellIn.Value() < 1 {
		change = 2
	} else {
		change = 1
	}

	a.quality.Update(change)
	a.sellIn.Tick()
}
