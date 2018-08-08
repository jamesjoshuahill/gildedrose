package inventory

type updater interface {
	update(*SellIn, *Quality)
}

type noopUpdater struct{}

func (noopUpdater) update(s *SellIn, q *Quality) {}

type standardUpdater struct {
	change, changePassedSellIn int
}

func (u standardUpdater) update(sellIn *SellIn, quality *Quality) {
	change := u.change
	if sellIn.Passed() {
		change = u.changePassedSellIn
	}
	quality.Update(change)

	sellIn.Decrement()
}

type backstagePassUpdater struct{}

func (backstagePassUpdater) update(sellIn *SellIn, quality *Quality) {
	days := sellIn.Days()
	change := 1
	if days <= 10 && days > 5 {
		change = 2
	}
	if days <= 5 && days > 0 {
		change = 3
	}
	if sellIn.Passed() {
		change = -quality.Value()
	}
	quality.Update(change)

	sellIn.Decrement()
}
