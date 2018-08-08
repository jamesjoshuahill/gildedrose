package inventory

const (
	agedBrie        = "Aged Brie"
	backstagePasses = "Backstage passes to a TAFKAL80ETC concert"
	conjured        = "Conjured Mana Cake"
	sulfuras        = "Sulfuras, Hand of Ragnaros"
)

type Item interface {
	Name() string
	SellIn() int
	Quality() int
	Update()
}

func NewItem(name string, sellInValue, qualityValue int) Item {
	sellIn := &sellIn{value: sellInValue}
	quality := &quality{value: qualityValue}

	switch name {
	case agedBrie:
		return item{name, sellIn, quality, standardUpdater{change: 1, changePassedSellIn: 2}}
	case conjured:
		return item{name, sellIn, quality, standardUpdater{change: -2, changePassedSellIn: -4}}
	case backstagePasses:
		return item{name, sellIn, quality, backstagePassUpdater{}}
	case sulfuras:
		return item{name, sellIn, quality, noopUpdater{}}
	default:
		return item{name, sellIn, quality, standardUpdater{change: -1, changePassedSellIn: -2}}
	}
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
