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
	{Name: "+5 Dexterity Vest", SellIn: 10, Quality: 20},
	{Name: "Aged Brie", SellIn: 2, Quality: 0},
	{Name: "Elixir of the Mongoose", SellIn: 5, Quality: 7},
	{Name: "Sulfuras, Hand of Ragnaros", SellIn: 0, Quality: 80},
	{Name: "Sulfuras, Hand of Ragnaros", SellIn: -1, Quality: 80},
	{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 15, Quality: 20},
	{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 10, Quality: 49},
	{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 5, Quality: 49},
	{Name: "Conjured Mana Cake", SellIn: 3, Quality: 6}, // Conjured items do not degrade twice as fast yet
}

func main() {
	daysArg := "0"
	if len(os.Args) == 2 {
		daysArg = os.Args[1]
	}

	days, err := strconv.Atoi(daysArg)
	if err != nil {
		log.Fatalf("error parsing days: %s\nusage: gildedrose [days]\n", err)
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
