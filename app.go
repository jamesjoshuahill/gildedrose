package gildedrose

type App struct {
	Items []Item
}

func NewApp(items []Item) *App {
	return &App{Items: items}
}

func (i *App) UpdateQuality() {
	for _, item := range i.Items {
		item.UpdateQuality()
	}
}
