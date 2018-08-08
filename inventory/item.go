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

func newItem(name string, sellIn, quality int) item {
	return item{name: name, sellIn: NewSellIn(sellIn), quality: NewQuality(quality)}
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

func newStandardUpdateItem(name string, sellIn, quality, change, changePassedSellIn int) *standardUpdateItem {
	return &standardUpdateItem{
		item:               newItem(name, sellIn, quality),
		change:             change,
		changePassedSellIn: changePassedSellIn,
	}
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
		return newStandardUpdateItem(name, sellIn, quality, -2, -4)
	}

	switch name {
	case agedBrie:
		return newStandardUpdateItem(name, sellIn, quality, 1, 2)
	case backstagePasses:
		return &BackstagePasses{item: newItem(name, sellIn, quality)}
	case sulfuras:
		return newItem(name, sellIn, quality)
	default:
		return newStandardUpdateItem(name, sellIn, quality, -1, -2)
	}
}
