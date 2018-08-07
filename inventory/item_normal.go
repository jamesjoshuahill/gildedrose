package inventory

type NormalItem struct {
	name    string
	sellIn  *SellIn
	quality *Quality
}

func (n NormalItem) Name() string {
	return n.name
}

func (n NormalItem) SellIn() int {
	return n.sellIn.Value()
}

func (n NormalItem) Quality() int {
	return n.quality.Value()
}

func (n *NormalItem) Update() {
	var change int
	if n.sellIn.Value() < 1 {
		change = -2
	} else {
		change = -1
	}

	n.quality.Update(change)
	n.sellIn.Tick()
}
