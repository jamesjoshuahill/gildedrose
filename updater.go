package gildedrose

type updater func(sellIn, quality int) (int, int)

func updateAgedBrie(sellIn, quality int) (int, int) {
	sellIn--

	quality = changeQuality(quality, 1)

	if sellIn < 0 {
		quality = changeQuality(quality, 1)
	}

	return sellIn, quality
}

func updateBackstagePasses(sellIn, quality int) (int, int) {
	sellIn--

	quality = changeQuality(quality, 1)

	if sellIn < 10 {
		quality = changeQuality(quality, 1)
	}

	if sellIn < 5 {
		quality = changeQuality(quality, 1)
	}

	if sellIn < 0 {
		quality = 0
	}

	return sellIn, quality
}

func updateConjured(sellIn, quality int) (int, int) {
	sellIn--

	quality = changeQuality(quality, -2)

	if sellIn < 0 {
		quality = changeQuality(quality, -2)
	}

	return sellIn, quality
}

func updateNormal(sellIn, quality int) (int, int) {
	sellIn--

	quality = changeQuality(quality, -1)

	if sellIn < 0 {
		quality = changeQuality(quality, -1)
	}

	return sellIn, quality
}

func updateSulfuras(sellIn, quality int) (int, int) {
	return sellIn, quality
}

func changeQuality(q, amount int) int {
	q = q + amount

	if q > 50 {
		q = 50
	}
	if q < 0 {
		q = 0
	}

	return q
}
