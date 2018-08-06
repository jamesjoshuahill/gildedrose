package inventory

type ItemBuilder struct{}

func (b ItemBuilder) Build(name string, sellIn, quality int) Item {
	return MagicItem{
		name:    name,
		sellIn:  sellIn,
		quality: quality,
	}
}
