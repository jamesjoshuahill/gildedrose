package inventory

import (
	"fmt"
)

const (
	agedBrie        = "Aged Brie"
	backstagePasses = "Backstage passes to a TAFKAL80ETC concert"
	sulfuras        = "Sulfuras, Hand of Ragnaros"
)

var items = []Item{
	{"+5 Dexterity Vest", 10, 20},
	{agedBrie, 2, 0},
	{"Elixir of the Mongoose", 5, 7},
	{sulfuras, 0, 80},
	{sulfuras, -1, 80},
	{backstagePasses, 15, 20},
	{backstagePasses, 10, 49},
	{backstagePasses, 5, 49},
	{"Conjured Mana Cake", 3, 6},
}

func Print() {
	for _, item := range items {
		fmt.Printf("%s, %d, %d\n", item.name, item.sellIn, item.quality)
	}
}

func Update() {
	for i := 0; i < len(items); i++ {

		if items[i].name != agedBrie && items[i].name != backstagePasses {
			if items[i].quality > 0 {
				if items[i].name != sulfuras {
					items[i].quality = items[i].quality - 1
				}
			}
		} else {
			if items[i].quality < 50 {
				items[i].quality = items[i].quality + 1
				if items[i].name == backstagePasses {
					if items[i].sellIn < 11 {
						if items[i].quality < 50 {
							items[i].quality = items[i].quality + 1
						}
					}
					if items[i].sellIn < 6 {
						if items[i].quality < 50 {
							items[i].quality = items[i].quality + 1
						}
					}
				}
			}
		}

		if items[i].name != sulfuras {
			items[i].sellIn = items[i].sellIn - 1
		}

		if items[i].sellIn < 0 {
			if items[i].name != agedBrie {
				if items[i].name != backstagePasses {
					if items[i].quality > 0 {
						if items[i].name != sulfuras {
							items[i].quality = items[i].quality - 1
						}
					}
				} else {
					items[i].quality -= items[i].quality
				}
			} else {
				if items[i].quality < 50 {
					items[i].quality = items[i].quality + 1
				}
			}
		}
	}
}
