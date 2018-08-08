package inventory

type SellIn struct {
	days int
}

func NewSellIn(value int) *SellIn {
	return &SellIn{days: value}
}

func (s SellIn) Days() int {
	return s.days
}

func (s SellIn) Passed() bool {
	return s.days <= 0
}

func (s *SellIn) Decrement() {
	s.days = s.days - 1
}
