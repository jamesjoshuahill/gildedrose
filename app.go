package gildedrose

type App struct {
	Items []Item
}

func NewApp(items []Item) *App {
	return &App{Items: items}
}

func (g *App) UpdateQuality() {
	for i := 0; i < len(g.Items); i++ {
		updated := NewUpdateableItem(g.Items[i]).Update()
		g.Items[i] = updated.Item
	}
}
