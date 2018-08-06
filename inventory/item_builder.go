package inventory

type ItemBuilder struct{}

func (b ItemBuilder) Build(name string, sellIn, quality int) Item {
	switch name {
	case agedBrie:
		return &AgedBrie{sellIn: NewSellIn(sellIn), quality: quality}
	case backstagePasses:
		return &BackstagePasses{sellIn: NewSellIn(sellIn), quality: quality}
	case sulfuras:
		return Sulfuras{sellIn: NewSellIn(sellIn), quality: quality}
	default:
		return &MagicItem{
			name:    name,
			sellIn:  NewSellIn(sellIn),
			quality: quality,
		}
	}
}
