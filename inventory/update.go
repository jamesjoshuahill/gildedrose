package inventory

const (
	agedBrie        = "Aged Brie"
	backstagePasses = "Backstage passes to a TAFKAL80ETC concert"
	sulfuras        = "Sulfuras, Hand of Ragnaros"
)

// Update changes the items in the inventory to reflect the passing of one day.
func (i *inventory) Update() {
	for n := 0; n < len(i.items); n++ {
		if i.items[n].name == sulfuras {
			continue
		}

		if i.items[n].name != agedBrie && i.items[n].name != backstagePasses {
			if i.items[n].quality > 0 {
				i.items[n].quality = i.items[n].quality - 1
			}
		} else {
			if i.items[n].quality < 50 {
				i.items[n].quality = i.items[n].quality + 1
				if i.items[n].name == backstagePasses {
					if i.items[n].sellIn < 11 {
						if i.items[n].quality < 50 {
							i.items[n].quality = i.items[n].quality + 1
						}
					}
					if i.items[n].sellIn < 6 {
						if i.items[n].quality < 50 {
							i.items[n].quality = i.items[n].quality + 1
						}
					}
				}
			}
		}

		i.items[n].sellIn -= 1

		if i.items[n].sellIn < 0 {
			if i.items[n].name != agedBrie {
				if i.items[n].name != backstagePasses {
					if i.items[n].quality > 0 {
						i.items[n].quality = i.items[n].quality - 1
					}
				} else {
					i.items[n].quality -= i.items[n].quality
				}
			} else {
				if i.items[n].quality < 50 {
					i.items[n].quality = i.items[n].quality + 1
				}
			}
		}
	}
}
