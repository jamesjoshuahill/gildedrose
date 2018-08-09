package gildedrose

type GildedRose struct {
	Items []MagicItem
}

func New(items []Item) *GildedRose {
	var magicItems []MagicItem
	for _, item := range items {
		magicItems = append(magicItems, NewMagicItem(item.Name, item.SellIn, item.Quality))
	}
	return &GildedRose{Items: magicItems}
}

func (i *GildedRose) UpdateQuality() {
	for _, item := range i.Items {
		item.UpdateQuality()
	}
}
