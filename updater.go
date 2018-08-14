package gildedrose

const (
	minQuality = 0
	maxQuality = 50
)

type updater func(sellIn, quality int) (int, int)

func noopUpdateFunc(sellIn, quality int) (int, int) {
	return sellIn, quality
}

func newStandardUpdateFunc(changeInDate, changeOutOfDate int) updater {
	return func(sellIn, quality int) (int, int) {
		sellIn = sellIn - 1

		change := changeInDate
		if sellIn < 0 {
			change = changeOutOfDate
		}

		return sellIn, normaliseQuality(quality, change)
	}
}

func backstagePassUpdateFunc(sellIn, quality int) (int, int) {
	sellIn = sellIn - 1

	change := 1
	if sellIn < 10 {
		change = 2
	}
	if sellIn < 5 {
		change = 3
	}
	if sellIn < 0 {
		change = -maxQuality
	}

	return sellIn, normaliseQuality(quality, change)
}

func normaliseQuality(value, change int) int {
	q := value + change

	if q < minQuality {
		q = minQuality
	}
	if q > maxQuality {
		q = maxQuality
	}

	return q
}
