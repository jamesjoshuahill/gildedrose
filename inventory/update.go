package inventory

import (
	"fmt"
)

const (
	agedBrie        = "Aged Brie"
	backstagePasses = "Backstage passes to a TAFKAL80ETC concert"
	sulfuras        = "Sulfuras, Hand of Ragnaros"
)

type inventory struct {
	items []Item
}

func New() *inventory {
	return &inventory{
		items: []Item{
			{name: "+5 Dexterity Vest", sellIn: 10, quality: 20},
			{name: agedBrie, sellIn: 2, quality: 0},
			{name: "Elixir of the Mongoose", sellIn: 5, quality: 7},
			{name: sulfuras, sellIn: 0, quality: 80},
			{name: sulfuras, sellIn: -1, quality: 80},
			{name: backstagePasses, sellIn: 15, quality: 20},
			{name: backstagePasses, sellIn: 10, quality: 49},
			{name: backstagePasses, sellIn: 5, quality: 49},
			{name: "Conjured Mana Cake", sellIn: 3, quality: 6},
		},
	}
}

func (i *inventory) List() (l string) {
	for _, item := range i.items {
		l += fmt.Sprintf("%s, %d, %d\n", item.name, item.sellIn, item.quality)
	}
	return
}

func (i *inventory) Update() {
	for n := 0; n < len(i.items); n++ {

		if i.items[n].name != agedBrie && i.items[n].name != backstagePasses {
			if i.items[n].quality > 0 {
				if i.items[n].name != sulfuras {
					i.items[n].quality = i.items[n].quality - 1
				}
			}
		} else {
			if i.items[n].quality < 50 {
				i.items[n].quality = i.items[n].quality + 1
				if i.items[n].name == backstagePasses {
					if i.items[n].sellIn < 11 {
						if i.items[n].quality < 50 {
							i.items[n].quality = i.items[n].quality + 1
						}
					}
					if i.items[n].sellIn < 6 {
						if i.items[n].quality < 50 {
							i.items[n].quality = i.items[n].quality + 1
						}
					}
				}
			}
		}

		if i.items[n].name != sulfuras {
			i.items[n].sellIn = i.items[n].sellIn - 1
		}

		if i.items[n].sellIn < 0 {
			if i.items[n].name != agedBrie {
				if i.items[n].name != backstagePasses {
					if i.items[n].quality > 0 {
						if i.items[n].name != sulfuras {
							i.items[n].quality = i.items[n].quality - 1
						}
					}
				} else {
					i.items[n].quality -= i.items[n].quality
				}
			} else {
				if i.items[n].quality < 50 {
					i.items[n].quality = i.items[n].quality + 1
				}
			}
		}
	}
}
