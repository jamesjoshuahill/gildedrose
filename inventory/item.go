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

func NewItem(name string, sellIn, quality int) Item {
	if strings.HasPrefix(name, "Conjured") {
		return &ConjuredItem{name: name, sellIn: NewSellIn(sellIn), quality: NewQuality(quality)}
	}

	switch name {
	case agedBrie:
		return &AgedBrie{sellIn: NewSellIn(sellIn), quality: NewQuality(quality)}
	case backstagePasses:
		return &BackstagePasses{sellIn: NewSellIn(sellIn), quality: NewQuality(quality)}
	case sulfuras:
		return Sulfuras{sellIn: NewSellIn(sellIn), quality: NewQuality(quality)}
	default:
		return &NormalItem{
			name:    name,
			sellIn:  NewSellIn(sellIn),
			quality: NewQuality(quality),
		}
	}
}
