package gildedrose

type App struct {
	Items []Item
}

func NewApp(items []Item) *App {
	return &App{Items: items}
}

func (g *App) UpdateQuality() {
	for i := 0; i < len(g.Items); i++ {
		switch g.Items[i].Name {
		case "Aged Brie":
			g.Items[i] = updateAgedBrie(g.Items[i])
		case "Backstage passes to a TAFKAL80ETC concert":
			g.Items[i] = updateBackstagePasses(g.Items[i])
		case "Sulfuras, Hand of Ragnaros":
			g.Items[i] = updateSulfuras(g.Items[i])
		default:
			g.Items[i] = updateNormal(g.Items[i])
		}
	}
}

func updateAgedBrie(item Item) Item {
	item.SellIn--

	item.Quality = incrementQuality(item.Quality)

	if item.SellIn < 0 {
		item.Quality = incrementQuality(item.Quality)
	}

	return item
}

func updateBackstagePasses(item Item) Item {
	item.SellIn--

	item.Quality = incrementQuality(item.Quality)

	if item.SellIn < 10 {
		item.Quality = incrementQuality(item.Quality)
	}

	if item.SellIn < 5 {
		item.Quality = incrementQuality(item.Quality)
	}

	if item.SellIn < 0 {
		item.Quality = 0
	}

	return item
}

func updateNormal(item Item) Item {
	item.SellIn--

	item.Quality = decrementQuality(item.Quality)

	if item.SellIn < 0 {
		item.Quality = decrementQuality(item.Quality)
	}

	return item
}

func updateSulfuras(item Item) Item {
	return item
}

func incrementQuality(q int) int {
	if q < 50 {
		return q + 1
	}
	return q
}

func decrementQuality(q int) int {
	if q > 0 {
		return q - 1
	}
	return q
}
