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

func (s SellIn) Passed() bool {
	return s.value <= 0
}

func (s *SellIn) Decrement() {
	s.value = s.value - 1
}
