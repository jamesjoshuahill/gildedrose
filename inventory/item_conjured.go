package inventory

type ConjuredItem struct {
	item
}

func (c *ConjuredItem) Update() {
	change := -2
	if c.sellIn.Passed() {
		change = -4
	}
	c.quality.Update(change)

	c.sellIn.Decrement()
}
