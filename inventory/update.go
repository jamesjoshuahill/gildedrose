package inventory

const (
	agedBrie        = "Aged Brie"
	backstagePasses = "Backstage passes to a TAFKAL80ETC concert"
	sulfuras        = "Sulfuras, Hand of Ragnaros"
)

// Update changes the items in the inventory to reflect the passing of one day.
func (i *inventory) Update() {
	for index, item := range i.items {
		if item.name == sulfuras {
			continue
		}

		i.items[index].sellIn--
		i.items[index].quality = updateQuality(item)
	}
}

func updateQuality(item Item) int {
	var change int

	switch item.name {
	case agedBrie:
		if item.sellIn < 1 {
			change = 2
		} else {
			change = 1
		}
	case backstagePasses:
		if item.sellIn > 10 {
			change = 1
		}
		if item.sellIn <= 10 && item.sellIn > 5 {
			change = 2
		}
		if item.sellIn <= 5 && item.sellIn > 0 {
			change = 3
		}
		if item.sellIn <= 0 {
			change = -item.quality
		}
	default:
		if item.sellIn < 1 {
			change = -2
		} else {
			change = -1
		}
	}

	return normaliseQuality(item.quality, change)
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
