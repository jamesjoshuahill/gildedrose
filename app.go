package gildedrose

type App struct {
	Items []Item
}

func NewApp(items []Item) *App {
	return &App{Items: items}
}

func (g *App) UpdateQuality() {
	for i := 0; i < len(g.Items); i++ {
		g.Items[i] = NewUpdateableItem(g.Items[i]).Update()
	}
}
