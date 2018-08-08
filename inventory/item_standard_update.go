package inventory

type standardUpdateItem struct {
	item
	change, changePassedSellIn int
}

func newStandardUpdateItem(name string, sellIn, quality, change, changePassedSellIn int) *standardUpdateItem {
	return &standardUpdateItem{
		item:               newItem(name, sellIn, quality),
		change:             change,
		changePassedSellIn: changePassedSellIn,
	}
}

func (q *standardUpdateItem) Update() {
	change := q.change
	if q.sellIn.Passed() {
		change = q.changePassedSellIn
	}
	q.quality.Update(change)

	q.sellIn.Decrement()
}
