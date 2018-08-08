package inventory

import "strings"

const (
	agedBrie        = "Aged Brie"
	backstagePasses = "Backstage passes to a TAFKAL80ETC concert"
	sulfuras        = "Sulfuras, Hand of Ragnaros"
)

type Item interface {
	Name() string
	SellIn() int
	Quality() int
	Update()
}

type item struct {
	name    string
	sellIn  *SellIn
	quality *Quality
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
	return
}

func NewItem(name string, sellIn, quality int) Item {
	if strings.HasPrefix(name, "Conjured") {
		return &ConjuredItem{item: item{name: name, sellIn: NewSellIn(sellIn), quality: NewQuality(quality)}}
	}

	switch name {
	case agedBrie:
		return &AgedBrie{item: item{name: name, sellIn: NewSellIn(sellIn), quality: NewQuality(quality)}}
	case backstagePasses:
		return &BackstagePasses{item: item{name: name, sellIn: NewSellIn(sellIn), quality: NewQuality(quality)}}
	case sulfuras:
		return item{name: name, sellIn: NewSellIn(sellIn), quality: NewQuality(quality)}
	default:
		return &NormalItem{item: item{name: name, sellIn: NewSellIn(sellIn), quality: NewQuality(quality)}}
	}
}
