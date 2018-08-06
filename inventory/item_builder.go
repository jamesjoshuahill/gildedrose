package inventory

type ItemBuilder struct{}

func (b ItemBuilder) Build(name string, sellIn, quality int) Item {
	switch name {
	case agedBrie:
		return &AgedBrie{sellIn: sellIn, quality: quality}
	case backstagePasses:
		return &BackstagePasses{sellIn: sellIn, quality: quality}
	case sulfuras:
		return Sulfuras{sellIn: sellIn, quality: quality}
	default:
		return &MagicItem{
			name:    name,
			sellIn:  sellIn,
			quality: quality,
		}
	}
}
