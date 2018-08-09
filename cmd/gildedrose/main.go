package main

import (
	"fmt"

	"github.com/jamesjoshuahill/gildedrose"
)

var items = []gildedrose.Item{
	{Name: "+5 Dexterity Vest", SellIn: 10, Quality: 20},
	{Name: "Aged Brie", SellIn: 2, Quality: 0},
	{Name: "Elixir of the Mongoose", SellIn: 5, Quality: 7},
	{Name: "Sulfuras, Hand of Ragnaros", SellIn: 0, Quality: 80},
	{Name: "Sulfuras, Hand of Ragnaros", SellIn: -1, Quality: 80},
	{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 15, Quality: 20},
	{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 10, Quality: 49},
	{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 5, Quality: 49},
	{Name: "Conjured Mana Cake", SellIn: 3, Quality: 10},
}

func main() {
	app := gildedrose.New(items)

	for day := 0; day <= 30; day++ {
		if day == 0 {
			fmt.Println("OMGHAI!")
		} else {
			fmt.Println()
		}
		fmt.Printf("-------- day %d --------\n", day)
		fmt.Printf("name, sellIn, quality\n")
		for _, item := range app.Items {
			fmt.Printf("%s, %d, %d\n", item.Name(), item.SellIn(), item.Quality())
		}
		app.UpdateQuality()
	}
}
