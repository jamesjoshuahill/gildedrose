package inventory

type Sulfuras struct {
	sellIn  *SellIn
	quality *Quality
}

func (s Sulfuras) Name() string {
	return sulfuras
}

func (s Sulfuras) SellIn() int {
	return s.sellIn.Value()
}

func (s Sulfuras) Quality() int {
	return s.quality.Value()
}

func (s Sulfuras) Update() {
	return
}
