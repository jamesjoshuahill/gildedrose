package inventory

const (
	AgedBrie        = "Aged Brie"
	BackstagePasses = "Backstage passes to a TAFKAL80ETC concert"
	Sulfuras        = "Sulfuras, Hand of Ragnaros"
)

type WrappedItem struct {
	Item
}

func (w WrappedItem) Update() WrappedItem {
	if w.Name == Sulfuras {
		return w
	}

	return WrappedItem{
		Item{Name: w.Name, SellIn: w.SellIn - 1, Quality: updateQuality(w)},
	}
}

func updateQuality(item WrappedItem) int {
	var change int

	switch item.Name {
	case AgedBrie:
		if item.SellIn < 1 {
			change = 2
		} else {
			change = 1
		}
	case BackstagePasses:
		if item.SellIn > 10 {
			change = 1
		}
		if item.SellIn <= 10 && item.SellIn > 5 {
			change = 2
		}
		if item.SellIn <= 5 && item.SellIn > 0 {
			change = 3
		}
		if item.SellIn <= 0 {
			change = -item.Quality
		}
	default:
		if item.SellIn < 1 {
			change = -2
		} else {
			change = -1
		}
	}

	return normaliseQuality(item.Quality, change)
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
