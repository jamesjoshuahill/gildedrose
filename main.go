package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/jamesjoshuahill/gildedrose/inventory"
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
	{Name: "Conjured Mana Cake", SellIn: 3, Quality: 6},
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

	var list []inventory.Item
	for _, item := range items {
		list = append(list, inventory.NewItem(item.Name, item.SellIn, item.Quality))
	}
	i := inventory.New(list)

	for day := 0; day <= days; day++ {
		if day == 0 {
			fmt.Println("OMGHAI!")
		} else {
			fmt.Println()
		}
		fmt.Printf("-------- day %d --------\n", day)
		fmt.Printf("name, sellIn, quality\n")
		for _, item := range i.List() {
			fmt.Printf("%s, %d, %d\n", item.Name(), item.SellIn(), item.Quality())
		}
		i.Update()
	}
}
