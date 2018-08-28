package gildedrose

type updateableItem struct {
	Item
	updater
}

func newUpdateableItem(item Item) updateableItem {
	switch item.Name {
	case "Aged Brie":
		return updateableItem{Item: item, updater: updateAgedBrie}
	case "Backstage passes to a TAFKAL80ETC concert":
		return updateableItem{Item: item, updater: updateBackstagePasses}
	case "Sulfuras, Hand of Ragnaros":
		return updateableItem{Item: item, updater: updateSulfuras}
	default:
		return updateableItem{Item: item, updater: updateNormal}
	}
}

func (u updateableItem) Update() Item {
	u.SellIn, u.Quality = u.updater(u.SellIn, u.Quality)
	return u.Item
}
