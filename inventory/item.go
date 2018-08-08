package inventory

type Item interface {
	Name() string
	SellIn() int
	Quality() int
	Update()
}

func NewItem(name string, sellInValue, qualityValue int) Item {
	var updater updater
	switch name {
	case "Aged Brie":
		updater = standardUpdater{change: 1, changePassedSellIn: 2}
	case "Conjured Mana Cake":
		updater = standardUpdater{change: -2, changePassedSellIn: -4}
	case "Backstage passes to a TAFKAL80ETC concert":
		updater = backstagePassUpdater{}
	case "Sulfuras, Hand of Ragnaros":
		updater = noopUpdater{}
	default:
		updater = standardUpdater{change: -1, changePassedSellIn: -2}
	}

	return item{name, &sellIn{value: sellInValue}, &quality{value: qualityValue}, updater}
}

type item struct {
	name    string
	sellIn  *sellIn
	quality *quality
	updater
}

func (b item) Name() string {
	return b.name
}

func (b item) SellIn() int {
	return b.sellIn.value
}

func (b item) Quality() int {
	return b.quality.value
}

func (b item) Update() {
	b.update(b.sellIn, b.quality)
}
