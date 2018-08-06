package inventory

const (
	AgedBrie        = "Aged Brie"
	BackstagePasses = "Backstage passes to a TAFKAL80ETC concert"
	Sulfuras        = "Sulfuras, Hand of Ragnaros"
)

// Update changes the items in the inventory to reflect the passing of one day.
func (i *inventory) Update() {
	for index, item := range i.items {
		if item.Name == Sulfuras {
			continue
		}

		i.items[index].SellIn--
		i.items[index].Quality = updateQuality(item)
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
