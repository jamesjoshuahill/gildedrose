package gildedrose

func NewMagicItem(name string, sellInValue, qualityValue int) MagicItem {
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

	return MagicItem{name, &sellIn{value: sellInValue}, &quality{value: qualityValue}, updater}
}

type MagicItem struct {
	name    string
	sellIn  *sellIn
	quality *quality
	updater
}

func (b MagicItem) Name() string {
	return b.name
}

func (b MagicItem) SellIn() int {
	return b.sellIn.value
}

func (b MagicItem) Quality() int {
	return b.quality.value
}

func (b *MagicItem) UpdateQuality() {
	b.update(b.sellIn, b.quality)
}
