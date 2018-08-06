package inventory

type Item interface {
	Name() string
	SellIn() int
	Quality() int
	Update() Item
}

type MagicItem struct {
	name            string
	sellIn, quality int
}

func (m MagicItem) Name() string {
	return m.name
}

func (m MagicItem) SellIn() int {
	return m.sellIn
}

func (m MagicItem) Quality() int {
	return m.quality
}

func (m MagicItem) Update() Item {
	if m.Name() == "Sulfuras, Hand of Ragnaros" {
		return m
	}

	return ItemBuilder{}.Build(m.Name(), m.SellIn()-1, updateQuality(m))
}

func updateQuality(item MagicItem) int {
	var change int

	switch item.Name() {
	case "Aged Brie":
		if item.SellIn() < 1 {
			change = 2
		} else {
			change = 1
		}
	case "Backstage passes to a TAFKAL80ETC concert":
		if item.SellIn() > 10 {
			change = 1
		}
		if item.SellIn() <= 10 && item.SellIn() > 5 {
			change = 2
		}
		if item.SellIn() <= 5 && item.SellIn() > 0 {
			change = 3
		}
		if item.SellIn() <= 0 {
			change = -item.Quality()
		}
	default:
		if item.SellIn() < 1 {
			change = -2
		} else {
			change = -1
		}
	}

	return normaliseQuality(item.Quality(), change)
}

func normaliseQuality(current int, change int) int {
	q := current + change

	if q < 0 {
		q = 0
	}
	if q > 50 {
		q = 50
	}

	return q
}
