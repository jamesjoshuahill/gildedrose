package gildedrose

type App struct {
	Items []Item
}

func NewApp(items []Item) *App {
	return &App{Items: items}
}

func (i *App) UpdateQuality() {
	for index, item := range i.Items {
		i.Items[index] = item.UpdateQuality()
	}
}
