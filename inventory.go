package gildedrose

type GildedRose struct {
	Items []MagicItem
}

func New(items []MagicItem) *GildedRose {
	return &GildedRose{Items: items}
}

func (i *GildedRose) Update() {
	for _, item := range i.Items {
		item.Update()
	}
}
