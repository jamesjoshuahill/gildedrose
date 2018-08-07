package inventory

type BackstagePasses struct {
	sellIn  *SellIn
	quality *Quality
}

func (b BackstagePasses) Name() string {
	return backstagePasses
}

func (b BackstagePasses) SellIn() int {
	return b.sellIn.Value()
}

func (b BackstagePasses) Quality() int {
	return b.quality.Value()
}

func (b *BackstagePasses) Update() {
	current := b.sellIn.Value()

	var change int
	if current > 10 {
		change = 1
	}
	if current <= 10 && current > 5 {
		change = 2
	}
	if current <= 5 && current > 0 {
		change = 3
	}
	if current <= 0 {
		change = -b.quality.Value()
	}

	b.quality.Update(change)
	b.sellIn.Tick()
}
