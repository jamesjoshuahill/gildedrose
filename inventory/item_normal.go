package inventory

type NormalItem struct {
	item
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
