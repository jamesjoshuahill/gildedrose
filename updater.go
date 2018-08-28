package gildedrose

type updater func(sellIn, quality int) (int, int)

func updateAgedBrie(sellIn, quality int) (int, int) {
	sellIn--

	quality = incrementQuality(quality)

	if sellIn < 0 {
		quality = incrementQuality(quality)
	}

	return sellIn, quality
}

func updateBackstagePasses(sellIn, quality int) (int, int) {
	sellIn--

	quality = incrementQuality(quality)

	if sellIn < 10 {
		quality = incrementQuality(quality)
	}

	if sellIn < 5 {
		quality = incrementQuality(quality)
	}

	if sellIn < 0 {
		quality = 0
	}

	return sellIn, quality
}

func updateConjured(sellIn, quality int) (int, int) {
	sellIn--

	quality = decrementQuality(decrementQuality(quality))

	if sellIn < 0 {
		quality = decrementQuality(decrementQuality(quality))
	}

	return sellIn, quality
}

func updateNormal(sellIn, quality int) (int, int) {
	sellIn--

	quality = decrementQuality(quality)

	if sellIn < 0 {
		quality = decrementQuality(quality)
	}

	return sellIn, quality
}

func updateSulfuras(sellIn, quality int) (int, int) {
	return sellIn, quality
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
