package inventory

type Quality struct {
	value int
}

func NewQuality(value int) *Quality {
	return &Quality{value: value}
}

func (q Quality) Value() int {
	return q.value
}

func (q *Quality) Update(amount int) {
	v := q.value + amount

	if v < 0 {
		v = 0
	}
	if v > 50 {
		v = 50
	}

	q.value = v
}
