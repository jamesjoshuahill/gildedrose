package inventory

type updater interface {
	update(*sellIn, *quality)
}

type noopUpdater struct{}

func (noopUpdater) update(s *sellIn, q *quality) {}

type standardUpdater struct {
	changeInDate, changeOutOfDate int
}

func (u standardUpdater) update(sellIn *sellIn, quality *quality) {
	sellIn.decrement()

	change := u.changeInDate
	if sellIn.lessThan(0) {
		change = u.changeOutOfDate
	}
	quality.update(change)
}

type backstagePassUpdater struct{}

func (backstagePassUpdater) update(sellIn *sellIn, quality *quality) {
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
