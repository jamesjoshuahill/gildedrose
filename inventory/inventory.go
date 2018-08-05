package inventory

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
