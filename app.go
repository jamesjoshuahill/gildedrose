package gildedrose

type App struct {
	Items []Item
}

func NewApp(items []Item) *App {
	return &App{Items: items}
}

func (g *App) UpdateQuality() {
	for i := 0; i < len(g.Items); i++ {
		if g.Items[i].Name == "Sulfuras, Hand of Ragnaros" {
			continue
		}

		if g.Items[i].Name == "Aged Brie" {
			g.Items[i].SellIn = g.Items[i].SellIn - 1

			if g.Items[i].Quality < 50 {
				g.Items[i].Quality = g.Items[i].Quality + 1
			}

			if g.Items[i].SellIn < 0 && g.Items[i].Quality < 50 {
				g.Items[i].Quality = g.Items[i].Quality + 1
			}

			continue
		}

		if g.Items[i].Name != "Backstage passes to a TAFKAL80ETC concert" {
			if g.Items[i].Quality > 0 {
				g.Items[i].Quality = g.Items[i].Quality - 1
			}
		} else {
			if g.Items[i].Quality < 50 {
				g.Items[i].Quality = g.Items[i].Quality + 1
				if g.Items[i].Name == "Backstage passes to a TAFKAL80ETC concert" {
					if g.Items[i].SellIn < 11 {
						if g.Items[i].Quality < 50 {
							g.Items[i].Quality = g.Items[i].Quality + 1
						}
					}
					if g.Items[i].SellIn < 6 {
						if g.Items[i].Quality < 50 {
							g.Items[i].Quality = g.Items[i].Quality + 1
						}
					}
				}
			}
		}

		g.Items[i].SellIn = g.Items[i].SellIn - 1

		if g.Items[i].SellIn < 0 {
			if g.Items[i].Name != "Backstage passes to a TAFKAL80ETC concert" {
				if g.Items[i].Quality > 0 {
					g.Items[i].Quality = g.Items[i].Quality - 1
				}
			} else {
				g.Items[i].Quality = g.Items[i].Quality - g.Items[i].Quality
			}
		}
	}
}
