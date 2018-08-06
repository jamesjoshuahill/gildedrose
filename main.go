package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

// Item defines an item in the inventory.
// Author: goblin in the corner. Do not touch!
type Item struct {
	Name            string
	SellIn, Quality int
}

// items is the list of items in the inventory.
// Author: goblin in the corner. Do not touch!
var items = []Item{
	{"+5 Dexterity Vest", 10, 20},
	{"Aged Brie", 2, 0},
	{"Elixir of the Mongoose", 5, 7},
	{"Sulfuras, Hand of Ragnaros", 0, 80},
	{"Sulfuras, Hand of Ragnaros", -1, 80},
	{"Backstage passes to a TAFKAL80ETC concert", 15, 20},
	{"Backstage passes to a TAFKAL80ETC concert", 10, 49},
	{"Backstage passes to a TAFKAL80ETC concert", 5, 49},
	{"Conjured Mana Cake", 3, 6},
}

func main() {
	daysArg := os.Args[1]
	if daysArg == "" {
		daysArg = "0"
	}

	days, err := strconv.Atoi(daysArg)
	if err != nil {
		log.Fatalln(err)
	}

	for day := 0; day <= days; day++ {
		if day == 0 {
			fmt.Println("OMGHAI!")
		} else {
			fmt.Println()
		}
		fmt.Printf("-------- day %d --------\n", day)
		fmt.Printf("name, sellIn, quality\n")
		for _, item := range items {
			fmt.Printf("%s, %d, %d\n", item.Name, item.SellIn, item.Quality)
		}
		GildedRose()
	}
}

func GildedRose() {
	for i := 0; i < len(items); i++ {

		if items[i].Name != "Aged Brie" && items[i].Name != "Backstage passes to a TAFKAL80ETC concert" {
			if items[i].Quality > 0 {
				if items[i].Name != "Sulfuras, Hand of Ragnaros" {
					items[i].Quality = items[i].Quality - 1
				}
			}
		} else {
			if items[i].Quality < 50 {
				items[i].Quality = items[i].Quality + 1
				if items[i].Name == "Backstage passes to a TAFKAL80ETC concert" {
					if items[i].SellIn < 11 {
						if items[i].Quality < 50 {
							items[i].Quality = items[i].Quality + 1
						}
					}
					if items[i].SellIn < 6 {
						if items[i].Quality < 50 {
							items[i].Quality = items[i].Quality + 1
						}
					}
				}
			}
		}

		if items[i].Name != "Sulfuras, Hand of Ragnaros" {
			items[i].SellIn = items[i].SellIn - 1
		}

		if items[i].SellIn < 0 {
			if items[i].Name != "Aged Brie" {
				if items[i].Name != "Backstage passes to a TAFKAL80ETC concert" {
					if items[i].Quality > 0 {
						if items[i].Name != "Sulfuras, Hand of Ragnaros" {
							items[i].Quality = items[i].Quality - 1
						}
					}
				} else {
					items[i].Quality = items[i].Quality - items[i].Quality
				}
			} else {
				if items[i].Quality < 50 {
					items[i].Quality = items[i].Quality + 1
				}
			}
		}
	}
}
