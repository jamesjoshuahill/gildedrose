package inventory

type Item struct {
	Name            string
	SellIn, Quality int
}

func (w Item) Update() Item {
	if w.Name == "Sulfuras, Hand of Ragnaros" {
		return w
	}

	return Item{
		Name:    w.Name,
		SellIn:  w.SellIn - 1,
		Quality: updateQuality(w),
	}
}

func updateQuality(item Item) int {
	var change int

	switch item.Name {
	case "Aged Brie":
		if item.SellIn < 1 {
			change = 2
		} else {
			change = 1
		}
	case "Backstage passes to a TAFKAL80ETC concert":
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
