package inventory

type NormalItem struct {
	item
}

func (n *NormalItem) Update() {
	change := -1
	if n.sellIn.Passed() {
		change = -2
	}
	n.quality.Update(change)

	n.sellIn.Decrement()
}
