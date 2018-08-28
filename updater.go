package gildedrose

type updater func(sellIn, quality int) (int, int)

func newStandardUpdater(changeBeforeSellBy, changeAfterSellBy int) updater {
	return func(sellIn, quality int) (int, int) {
		sellIn--

		change := changeBeforeSellBy

		if sellIn < 0 {
			change = changeAfterSellBy
		}

		return sellIn, changeQuality(quality, change)
	}
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
