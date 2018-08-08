package inventory

type ConjuredItem struct {
	item
}

func (c *ConjuredItem) Update() {
	var change int
	if c.sellIn.Value() < 1 {
		change = -4
	} else {
		change = -2
	}

	c.quality.Update(change)
	c.sellIn.Tick()
}
