package gildedrose

type GildedRose struct {
	Items []Item
}

func New(items []Item) *GildedRose {
	return &GildedRose{Items: items}
}

func (i *GildedRose) UpdateQuality() {
	for _, item := range i.Items {
		item.UpdateQuality()
	}
}
