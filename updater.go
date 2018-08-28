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

func backstagePassesUpdater(sellIn, quality int) (int, int) {
	sellIn--

	change := 1
	if sellIn < 10 {
		change = 2
	}
	if sellIn < 5 {
		change = 3
	}
	if sellIn < 0 {
		change = -quality
	}

	return sellIn, changeQuality(quality, change)
}

func noopUpdater(sellIn, quality int) (int, int) {
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
