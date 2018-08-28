package gildedrose

type updater func(item Item) Item

func updateAgedBrie(item Item) Item {
	item.SellIn--

	item.Quality = incrementQuality(item.Quality)

	if item.SellIn < 0 {
		item.Quality = incrementQuality(item.Quality)
	}

	return item
}

func updateBackstagePasses(item Item) Item {
	item.SellIn--

	item.Quality = incrementQuality(item.Quality)

	if item.SellIn < 10 {
		item.Quality = incrementQuality(item.Quality)
	}

	if item.SellIn < 5 {
		item.Quality = incrementQuality(item.Quality)
	}

	if item.SellIn < 0 {
		item.Quality = 0
	}

	return item
}

func updateNormal(item Item) Item {
	item.SellIn--

	item.Quality = decrementQuality(item.Quality)

	if item.SellIn < 0 {
		item.Quality = decrementQuality(item.Quality)
	}

	return item
}

func updateSulfuras(item Item) Item {
	return item
}

func incrementQuality(q int) int {
	if q < 50 {
		return q + 1
	}
	return q
}

func decrementQuality(q int) int {
	if q > 0 {
		return q - 1
	}
	return q
}
