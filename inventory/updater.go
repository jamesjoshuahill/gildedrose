package inventory

type updater interface {
	update(*sellIn, *quality)
}

type noopUpdater struct{}

func (noopUpdater) update(s *sellIn, q *quality) {}

type standardUpdater struct {
	change, changePassedSellIn int
}

func (u standardUpdater) update(sellIn *sellIn, quality *quality) {
	change := u.change
	if sellIn.passed() {
		change = u.changePassedSellIn
	}
	quality.update(change)

	sellIn.decrement()
}

type backstagePassUpdater struct{}

func (backstagePassUpdater) update(sellIn *sellIn, quality *quality) {
	change := 1
	if sellIn.between(6, 10) {
		change = 2
	}
	if sellIn.between(1, 5) {
		change = 3
	}
	if sellIn.passed() {
		change = -quality.value
	}
	quality.update(change)

	sellIn.decrement()
}
