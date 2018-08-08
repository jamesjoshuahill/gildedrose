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
	change := u.changeInDate
	if sellIn.lessThan(1) {
		change = u.changeOutOfDate
	}
	quality.update(change)

	sellIn.decrement()
}

type backstagePassUpdater struct{}

func (backstagePassUpdater) update(sellIn *sellIn, quality *quality) {
	change := 1
	if sellIn.lessThan(11) {
		change = 2
	}
	if sellIn.lessThan(6) {
		change = 3
	}
	if sellIn.lessThan(1) {
		change = -maxQuality
	}
	quality.update(change)

	sellIn.decrement()
}
