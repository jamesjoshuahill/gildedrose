package inventory

type ItemBuilder struct{}

func (b ItemBuilder) Build(name string, sellIn, quality int) Item {
	if name == sulfuras {
		return Sulfuras{sellIn: sellIn, quality: quality}
	}

	if name == agedBrie {
		return AgedBrie{sellIn: sellIn, quality: quality}
	}

	return MagicItem{
		name:    name,
		sellIn:  sellIn,
		quality: quality,
	}
}
