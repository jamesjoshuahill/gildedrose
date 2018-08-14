package gildedrose

type updater func(*sellIn, *quality)

func noopUpdateFunc(_ *sellIn, _ *quality) {}

func newStandardUpdateFunc(changeInDate, changeOutOfDate int) updater {
	return func(sellIn *sellIn, quality *quality) {
		sellIn.decrement()

		change := changeInDate
		if sellIn.lessThan(0) {
			change = changeOutOfDate
		}
		quality.update(change)
	}
}

func backstagePassUpdateFunc(sellIn *sellIn, quality *quality) {
	sellIn.decrement()

	change := 1
	if sellIn.lessThan(10) {
		change = 2
	}
	if sellIn.lessThan(5) {
		change = 3
	}
	if sellIn.lessThan(0) {
		change = -maxQuality
	}
	quality.update(change)
}
