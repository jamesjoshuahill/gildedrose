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

		g.Items[i].SellIn = g.Items[i].SellIn - 1

		if g.Items[i].Name == "Aged Brie" {
			g.Items[i].Quality = incrementQuality(g.Items[i].Quality)

			if g.Items[i].SellIn < 0 {
				g.Items[i].Quality = incrementQuality(g.Items[i].Quality)
			}

			continue
		}

		if g.Items[i].Name == "Backstage passes to a TAFKAL80ETC concert" {
			g.Items[i].Quality = incrementQuality(g.Items[i].Quality)

			if g.Items[i].SellIn < 10 {
				g.Items[i].Quality = incrementQuality(g.Items[i].Quality)
			}

			if g.Items[i].SellIn < 5 {
				g.Items[i].Quality = incrementQuality(g.Items[i].Quality)
			}

			if g.Items[i].SellIn < 0 {
				g.Items[i].Quality = g.Items[i].Quality - g.Items[i].Quality
			}

			continue
		}

		if g.Items[i].Quality > 0 {
			g.Items[i].Quality = g.Items[i].Quality - 1
		}

		if g.Items[i].SellIn < 0 {
			if g.Items[i].Quality > 0 {
				g.Items[i].Quality = g.Items[i].Quality - 1
			}
		}
	}
}

func incrementQuality(q int) int {
	if q < 50 {
		return q + 1
	}
	return q
}
