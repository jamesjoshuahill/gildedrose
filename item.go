package gildedrose

type Item struct {
	name    string
	sellIn  int
	quality int
	updater updater
}

func NewItem(name string, sellIn, quality int) *Item {
	var updater updater
	switch name {
	case "Aged Brie":
		updater = newStandardUpdateFunc(1, 2)
	case "Conjured Mana Cake":
		updater = newStandardUpdateFunc(-2, -4)
	case "Backstage passes to a TAFKAL80ETC concert":
		updater = backstagePassUpdateFunc
	case "Sulfuras, Hand of Ragnaros":
		updater = noopUpdateFunc
	default:
		updater = newStandardUpdateFunc(-1, -2)
	}

	return &Item{
		name:    name,
		sellIn:  sellIn,
		quality: quality,
		updater: updater,
	}
}

func (b Item) Name() string {
	return b.name
}

func (b Item) SellIn() int {
	return b.sellIn
}

func (b Item) Quality() int {
	return b.quality
}

func (b *Item) UpdateQuality() {
	b.sellIn, b.quality = b.updater(b.sellIn, b.quality)
}
