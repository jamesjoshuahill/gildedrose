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
	s := NewSellIn(sellIn)
	q := NewQuality(quality)

	if strings.HasPrefix(name, conjured) {
		return item{name, s, q, standardUpdater{change: -2, changePassedSellIn: -4}}
	}

	switch name {
	case agedBrie:
		return item{name, s, q, standardUpdater{change: 1, changePassedSellIn: 2}}
	case backstagePasses:
		return item{name, s, q, backstagePassUpdater{}}
	case sulfuras:
		return item{name, s, q, noopUpdater{}}
	default:
		return item{name, s, q, standardUpdater{change: -1, changePassedSellIn: -2}}
	}
}

type item struct {
	name    string
	sellIn  *SellIn
	quality *Quality
	updater
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
