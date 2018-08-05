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
			if item.sellIn < 1 {
				item.quality++
			}
			item.quality++
		case backstagePasses:
			if item.sellIn > 10 {
				item.quality++
			}
			if item.sellIn <= 10 && item.sellIn > 5 {
				item.quality += 2
			}
			if item.sellIn <= 5 && item.sellIn > 0 {
				item.quality += 3
			}
			if item.sellIn <= 0 {
				item.quality = 0
			}
		case sulfuras:
			continue
		default:
			if item.sellIn < 1 {
				item.quality--
			}
			item.quality--
		}

		if item.quality < 0 {
			item.quality = 0
		}
		if item.quality > 50 {
			item.quality = 50
		}

		item.sellIn--

		i.items[index] = item
	}
}
