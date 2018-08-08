package inventory

type sellIn struct {
	value int
}

func (s sellIn) passed() bool {
	return s.value <= 0
}

func (s sellIn) lessThan(days int) bool {
	return s.value < days
}

func (s *sellIn) decrement() {
	s.value = s.value - 1
}
