package inventory

type ConjuredItem struct {
	item
}

func (c *ConjuredItem) Update() {
	var change int
	if c.sellIn.Passed() {
		change = -4
	} else {
		change = -2
	}

	c.quality.Update(change)
	c.sellIn.Decrement()
}
