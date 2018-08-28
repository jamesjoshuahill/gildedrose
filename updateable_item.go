package gildedrose

type updateableItem struct {
	Item
	updater
}

func NewUpdateableItem(item Item) updateableItem {
	switch item.Name {
	case "Aged Brie":
		return updateableItem{Item: item, updater: newStandardUpdater(1, 2)}
	case "Backstage passes to a TAFKAL80ETC concert":
		return updateableItem{Item: item, updater: backstagePassesUpdater}
	case "Conjured Mana Cake":
		return updateableItem{Item: item, updater: newStandardUpdater(-2, -4)}
	case "Sulfuras, Hand of Ragnaros":
		return updateableItem{Item: item, updater: noopUpdater}
	default:
		return updateableItem{Item: item, updater: newStandardUpdater(-1, -2)}
	}
}

func (u updateableItem) Update() updateableItem {
	u.SellIn, u.Quality = u.updater(u.SellIn, u.Quality)
	return u
}
