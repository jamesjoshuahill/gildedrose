package inventory

// Item defines an item in the inventory. Do not touch!
type Item struct {
	name            string
	sellIn, quality int
}

// items is the list of items in the inventory. Do not touch!
var items = []Item{
	{name: "+5 Dexterity Vest", sellIn: 10, quality: 20},
	{name: agedBrie, sellIn: 2, quality: 0},
	{name: "Elixir of the Mongoose", sellIn: 5, quality: 7},
	{name: sulfuras, sellIn: 0, quality: 80},
	{name: sulfuras, sellIn: -1, quality: 80},
	{name: backstagePasses, sellIn: 15, quality: 20},
	{name: backstagePasses, sellIn: 10, quality: 49},
	{name: backstagePasses, sellIn: 5, quality: 49},
	{name: "Conjured Mana Cake", sellIn: 3, quality: 6},
}
