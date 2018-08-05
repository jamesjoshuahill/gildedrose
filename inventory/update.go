package inventory

const (
	agedBrie        = "Aged Brie"
	backstagePasses = "Backstage passes to a TAFKAL80ETC concert"
	sulfuras        = "Sulfuras, Hand of Ragnaros"
)

// Update changes the items in the inventory to reflect the passing of one day.
func (i *inventory) Update() {
	for index, item := range i.items {
		switch item.name {
		case agedBrie:
			if item.quality < 50 {
				item.quality++
			}
		case backstagePasses:
			item.quality++

			if item.sellIn < 11 {
				item.quality++
			}
			if item.sellIn < 6 {
				item.quality++
			}
			if item.quality > 50 {
				item.quality = 50
			}
		case sulfuras:
			continue
		default:
			if item.quality > 0 {
				item.quality--
			}
		}

		item.sellIn--

		if item.sellIn < 0 {
			switch item.name {
			case agedBrie:
				if item.quality < 50 {
					item.quality++
				}
			case backstagePasses:
				item.quality = 0
			default:
				if item.quality > 0 {
					item.quality--
				}
			}
		}

		i.items[index] = item
	}
}
