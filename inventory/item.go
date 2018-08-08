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

type standardUpdateItem struct {
	item
	change, changePassedSellIn int
}

func (q *standardUpdateItem) Update() {
	change := q.change
	if q.sellIn.Passed() {
		change = q.changePassedSellIn
	}
	q.quality.Update(change)

	q.sellIn.Decrement()
}

func NewItem(name string, sellIn, quality int) Item {
	if strings.HasPrefix(name, "Conjured") {
		return &standardUpdateItem{
			item:               item{name: name, sellIn: NewSellIn(sellIn), quality: NewQuality(quality)},
			change:             -2,
			changePassedSellIn: -4,
		}
	}

	switch name {
	case agedBrie:
		return &standardUpdateItem{
			item:               item{name: name, sellIn: NewSellIn(sellIn), quality: NewQuality(quality)},
			change:             1,
			changePassedSellIn: 2,
		}
	case backstagePasses:
		return &BackstagePasses{item: item{name: name, sellIn: NewSellIn(sellIn), quality: NewQuality(quality)}}
	case sulfuras:
		return item{name: name, sellIn: NewSellIn(sellIn), quality: NewQuality(quality)}
	default:
		return &standardUpdateItem{
			item:               item{name: name, sellIn: NewSellIn(sellIn), quality: NewQuality(quality)},
			change:             -1,
			changePassedSellIn: -2,
		}
	}
}
