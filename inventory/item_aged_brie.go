package inventory

type AgedBrie struct {
	sellIn  *SellIn
	quality *Quality
}

func (a AgedBrie) Name() string {
	return agedBrie
}

func (a AgedBrie) SellIn() int {
	return a.sellIn.Value()
}

func (a AgedBrie) Quality() int {
	return a.quality.Value()
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
