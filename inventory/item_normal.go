package inventory

type NormalItem struct {
	item
}

func (n *NormalItem) Update() {
	var change int
	if n.sellIn.Passed() {
		change = -2
	} else {
		change = -1
	}

	n.quality.Update(change)
	n.sellIn.Decrement()
}
