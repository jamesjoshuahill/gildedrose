package main

import (
	"fmt"

	"github.com/jamesjoshuahill/gildedrose"
)

func main() {
	fmt.Println("OMGHAI!")

	items := []gildedrose.Item{
		gildedrose.NewItem("+5 Dexterity Vest", 10, 20),
		gildedrose.NewItem("Aged Brie", 2, 0),
		gildedrose.NewItem("Elixir of the Mongoose", 5, 7),
		gildedrose.NewItem("Sulfuras, Hand of Ragnaros", 0, 80),
		gildedrose.NewItem("Sulfuras, Hand of Ragnaros", -1, 80),
		gildedrose.NewItem("Backstage passes to a TAFKAL80ETC concert", 15, 20),
		gildedrose.NewItem("Backstage passes to a TAFKAL80ETC concert", 10, 49),
		gildedrose.NewItem("Backstage passes to a TAFKAL80ETC concert", 5, 49),
		gildedrose.NewItem("Conjured Mana Cake", 3, 10),
	}

	app := gildedrose.New(items)

	for day := 0; day <= 30; day++ {
		fmt.Printf("-------- day %d --------\n", day)
		fmt.Printf("name, sellIn, quality\n")
		for _, item := range app.Items {
			fmt.Printf("%s, %d, %d\n", item.Name(), item.SellIn(), item.Quality())
		}
		if day < 30 {
			fmt.Println()
		}
		app.UpdateQuality()
	}
}
