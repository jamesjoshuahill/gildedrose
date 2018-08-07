package inventory

type SellIn struct {
	value int
}

func NewSellIn(value int) *SellIn {
	return &SellIn{value: value}
}

func (s SellIn) Value() int {
	return s.value
}

func (s *SellIn) Tick() {
	s.value = s.value - 1
}
