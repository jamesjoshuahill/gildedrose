package inventory

type ConjuredItem struct {
	name    string
	sellIn  *SellIn
	quality *Quality
}

func (c ConjuredItem) Name() string {
	return c.name
}

func (c ConjuredItem) SellIn() int {
	return c.sellIn.Value()
}

func (c ConjuredItem) Quality() int {
	return c.quality.Value()
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
