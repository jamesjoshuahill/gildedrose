package inventory

const (
	agedBrie        = "Aged Brie"
	backstagePasses = "Backstage passes to a TAFKAL80ETC concert"
	sulfuras        = "Sulfuras, Hand of Ragnaros"
)

// Update changes the items in the inventory to reflect the passing of one day.
func (i *inventory) Update() {
	for index, item := range i.items {
		var qualityChange int

		switch item.name {
		case agedBrie:
			if item.sellIn < 1 {
				qualityChange = 2
			} else {
				qualityChange = 1
			}
		case backstagePasses:
			if item.sellIn > 10 {
				qualityChange = 1
			}
			if item.sellIn <= 10 && item.sellIn > 5 {
				qualityChange = 2
			}
			if item.sellIn <= 5 && item.sellIn > 0 {
				qualityChange = 3
			}
			if item.sellIn <= 0 {
				qualityChange = -item.quality
			}
		case sulfuras:
			continue
		default:
			if item.sellIn < 1 {
				qualityChange = -2
			} else {
				qualityChange = -1
			}
		}

		i.items[index].sellIn--
		i.items[index].quality = normaliseQuality(item.quality, qualityChange)
	}
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
