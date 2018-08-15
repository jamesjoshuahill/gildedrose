package gildedrose

type App struct {
	items []Item
}

func NewApp(items []Item) *App {
	return &App{items: items}
}

func (i *App) Items() []Item {
	return i.items
}

func (i *App) UpdateQuality() {
	for index, item := range i.items {
		i.items[index] = item.UpdateQuality()
	}
}
