package inventory

import "strings"

const (
	agedBrie        = "Aged Brie"
	backstagePasses = "Backstage passes to a TAFKAL80ETC concert"
	conjured        = "Conjured"
	sulfuras        = "Sulfuras, Hand of Ragnaros"
)

type Item interface {
	Name() string
	SellIn() int
	Quality() int
	Update()
}

func NewItem(name string, sellIn, quality int) Item {
	if strings.HasPrefix(name, conjured) {
		return newItem(name, sellIn, quality, standardUpdater{change: -2, changePassedSellIn: -4})
	}

	switch name {
	case agedBrie:
		return newItem(name, sellIn, quality, standardUpdater{change: 1, changePassedSellIn: 2})
	case backstagePasses:
		return newItem(name, sellIn, quality, backstagePassUpdater{})
	case sulfuras:
		return newItem(name, sellIn, quality, noopUpdater{})
	default:
		return newItem(name, sellIn, quality, standardUpdater{change: -1, changePassedSellIn: -2})
	}
}

type item struct {
	name    string
	sellIn  *SellIn
	quality *Quality
	updater
}

func newItem(name string, sellIn, quality int, updater updater) item {
	return item{name: name, sellIn: NewSellIn(sellIn), quality: NewQuality(quality), updater: updater}
}

func (b item) Name() string {
	return b.name
}

func (b item) SellIn() int {
	return b.sellIn.Days()
}

func (b item) Quality() int {
	return b.quality.Value()
}

func (b item) Update() {
	b.update(b.sellIn, b.quality)
}
