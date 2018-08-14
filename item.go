package gildedrose

type Item struct {
	name    string
	sellIn  *sellIn
	quality *quality
	updater updater
}

func NewItem(name string, sellInValue, qualityValue int) Item {
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

	return Item{
		name:    name,
		sellIn:  &sellIn{value: sellInValue},
		quality: &quality{value: qualityValue},
		updater: updater,
	}
}

func (b Item) Name() string {
	return b.name
}

func (b Item) SellIn() int {
	return b.sellIn.value
}

func (b Item) Quality() int {
	return b.quality.value
}

func (b *Item) UpdateQuality() {
	b.updater(b.sellIn, b.quality)
}
