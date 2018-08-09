package gildedrose

type Item struct {
	name    string
	sellIn  *sellIn
	quality *quality
	updater
}

func NewItem(name string, sellInValue, qualityValue int) Item {
	var updater updater
	switch name {
	case "Aged Brie":
		updater = standardUpdater{changeInDate: 1, changeOutOfDate: 2}
	case "Conjured Mana Cake":
		updater = standardUpdater{changeInDate: -2, changeOutOfDate: -4}
	case "Backstage passes to a TAFKAL80ETC concert":
		updater = backstagePassUpdater{}
	case "Sulfuras, Hand of Ragnaros":
		updater = noopUpdater{}
	default:
		updater = standardUpdater{changeInDate: -1, changeOutOfDate: -2}
	}

	return Item{name, &sellIn{value: sellInValue}, &quality{value: qualityValue}, updater}
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
	b.update(b.sellIn, b.quality)
}
